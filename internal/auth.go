package internal

import (
	"context"
	"fmt"

	"github.com/busy-cloud/boat/api"
	"github.com/busy-cloud/boat/db"
	"github.com/busy-cloud/boat/web"
	"github.com/gin-gonic/gin"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram"
)

// WechatLoginRequest 微信登录请求
type WechatLoginRequest struct {
	Code string `json:"code" binding:"required"`
}

// WechatAuth 微信登录处理函数
func WechatAuth(ctx *gin.Context, mp *miniProgram.MiniProgram) {
	var req WechatLoginRequest

	// 从请求体中获取 code
	if err := ctx.ShouldBindJSON(&req); err != nil {
		api.Error(ctx, err)
		return
	}

	// 使用 code 换取 openid 和 session_key
	sessionResp, err := mp.Auth.Session(context.Background(), req.Code)
	if err != nil {
		api.Error(ctx, fmt.Errorf("微信登录失败: %v", err))
		return
	}

	// 检查是否成功获取到 openid
	if sessionResp.OpenID == "" {
		api.Fail(ctx, "获取 openid 失败")
		return
	}

	// 根据 openid 查找或创建用户
	var user WechatUser
	has, err := db.Engine().Where("openid=?", sessionResp.OpenID).Get(&user)
	if err != nil {
		api.Error(ctx, fmt.Errorf("查询用户失败: %v", err))
		return
	}

	// 如果用户不存在，创建新用户
	if !has {
		// 生成用户ID（使用 openid 的前16位，或者使用其他生成策略）
		userId := generateUserId(sessionResp.OpenID)
		
		user = WechatUser{
			Id:      userId,
			OpenID:  sessionResp.OpenID,
			UnionID: sessionResp.UnionID,
			Name:    "微信用户", // 默认名称，后续可以通过用户信息接口更新
		}

		// 插入新用户
		_, err = db.Engine().InsertOne(&user)
		if err != nil {
			api.Error(ctx, fmt.Errorf("创建用户失败: %v", err))
			return
		}
	}

	// 检查用户是否被禁用
	if user.Disabled {
		api.Fail(ctx, "用户已被禁用")
		return
	}

	// 更新 unionid（如果之前没有）
	if sessionResp.UnionID != "" && user.UnionID == "" {
		user.UnionID = sessionResp.UnionID
		_, err = db.Engine().ID(user.Id).Cols("unionid").Update(&user)
		if err != nil {
			// 更新失败不影响登录，只记录错误
			fmt.Printf("更新 unionid 失败: %v\n", err)
		}
	}

	// 生成 JWT token
	token, err := web.JwtGenerate(user.Id, false, "")
	if err != nil {
		api.Error(ctx, fmt.Errorf("生成 token 失败: %v", err))
		return
	}

	// 返回登录结果（格式与普通登录接口保持一致）
	api.OK(ctx, gin.H{
		"token": token,
		"user":  &user,
	})
}

// generateUserId 生成用户ID
// 使用 openid 的完整值作为用户ID的一部分，确保唯一性
func generateUserId(openid string) string {
	// 使用 "wx_" 前缀 + openid 作为用户ID
	// openid 通常是28位字符串，加上前缀后总长度在合理范围内
	return "wx_" + openid
}

