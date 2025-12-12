package weixin

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram"
	"github.com/busy-cloud/boat/boot"
	"github.com/busy-cloud/boat/db"
	"github.com/busy-cloud/weixin/internal"
)

func init() {
	boot.Register("weixin", &boot.Task{
		Startup:  startup,
		Shutdown: nil,
		Depends:  nil,
	})
}

var mp *miniProgram.MiniProgram

func startup() (err error) {
	// 初始化微信小程序配置
	cfg := miniProgram.UserConfig{
		AppID:  "wx0e77be5ee5eca13a",
		Secret: "a662eb919cc6d83a9b2e5f65c3742cdc",
	}

	mp, err = miniProgram.NewMiniProgram(&cfg)
	if err != nil {
		return err
	}

	// 同步数据库表结构
	err = db.Engine().Sync2(new(internal.WechatUser))
	if err != nil {
		return err
	}

	return nil
}
