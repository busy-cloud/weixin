package weixin

import (
	"embed"
	"encoding/json"
	"time"

	"github.com/busy-cloud/boat/apps"
	"github.com/busy-cloud/boat/log"
	"github.com/busy-cloud/boat/store"
)

type User struct {
	Id        string    `json:"id" xorm:"pk"`
	OpenId    string    `json:"openid,omitempty" xorm:"'openid'"`
	UnionId   string    `json:"unionid,omitempty" xorm:"'unionid'"`
	Name      string    `json:"name,omitempty"`
	Avatar    string    `json:"avatar,omitempty"`
	Cellphone string    `json:"cellphone,omitempty"`
	Admin     bool      `json:"admin,omitempty"`
	Disabled  bool      `json:"disabled,omitempty"`
	Created   time.Time `json:"created,omitempty" xorm:"created"`
	Updated   time.Time `json:"updated,omitempty" xorm:"updated"`
}

//go:embed tables
var tables embed.FS

//go:embed manifest.json
var manifest []byte

func init() {
	//注册为内部插件
	var a apps.App
	err := json.Unmarshal(manifest, &a)
	if err != nil {
		log.Fatal(err)
	}
	apps.Register(&a)

	//注册资源
	//a.AssetsFS = store.PrefixFS(&assets, "assets")
	//a.PagesFS = store.PrefixFS(&pages, "pages")
	a.TablesFS = store.PrefixFS(&tables, "tables")
}
