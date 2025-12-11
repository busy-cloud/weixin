package weixin

import (
	"context"

	"github.com/busy-cloud/boat/api"
	"github.com/gin-gonic/gin"
)

func init() {

	api.RegisterUnAuthorized("POST", "weixin/auth", func(ctx *gin.Context) {
		resp, err := mp.Auth.Session(context.Background(), ctx.Param("code"))
		if err != nil {
			api.Error(ctx, err)
			return
		}

		//resp.OpenID
		//TODO 自动登录

		//自动获取手机号
		resp2, err := mp.PhoneNumber.GetUserPhoneNumber(context.Background(), ctx.Param("code"))
		if err != nil {
			api.Error(ctx, err)
			return
		}

		api.OK(ctx, resp)
	})

	api.RegisterUnAuthorized("GET", "weixin/phone", func(ctx *gin.Context) {

	})

}
