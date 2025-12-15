package weixin

import (
	"github.com/busy-cloud/boat/config"
)

const MODULE = "weixin"

func init() {
	config.SetDefault(MODULE, "appid", "")
	config.SetDefault(MODULE, "secret", "")
}
