package weixin

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram"
	"github.com/busy-cloud/boat/boot"
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
	cfg := miniProgram.UserConfig{
		AppID:  "",
		Secret: "",
	}

	mp, err = miniProgram.NewMiniProgram(&cfg)
	return
}
