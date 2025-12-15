package weixin

import (
	"github.com/busy-cloud/boat/config"
	"github.com/busy-cloud/boat/smart"
)

func init() {
	config.Register(MODULE, &config.Form{
		Title:  "微信配置",
		Module: MODULE,
		Fields: []smart.Field{
			{Key: "appid", Label: "AppID", Type: "text"},
			{Key: "secret", Label: "AppSecret", Type: "text"},
		},
	})
}
