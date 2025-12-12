package weixin

import (
	"context"

	"github.com/busy-cloud/boat/api"
	"github.com/busy-cloud/weixin/internal"
	_ "github.com/busy-cloud/weixin/internal"
	"github.com/gin-gonic/gin"
)

func init() {
	// 微信登录接口
	api.RegisterUnAuthorized("POST", "weixin/auth", func(ctx *gin.Context) {
		// 获取已初始化的 mp 实例
		if mp == nil {
			api.Error(ctx, gin.H{"error": "微信小程序未初始化"})
			return
		}
		internal.WechatAuth(ctx, mp)
	})

	// 获取手机号接口
	api.RegisterUnAuthorized("GET", "weixin/phone", func(ctx *gin.Context) {
		// 获取已初始化的 mp 实例
		if mp == nil {
			api.Error(ctx, gin.H{"error": "微信小程序未初始化"})
			return
		}

		code := ctx.Query("code")
		if code == "" {
			api.Fail(ctx, "code 参数不能为空")
			return
		}

		// 使用 code 获取用户手机号
		phoneResp, err := mp.PhoneNumber.GetUserPhoneNumber(context.Background(), code)
		if err != nil {
			api.Error(ctx, err)
			return
		}

		api.OK(ctx, phoneResp)
	})
}
