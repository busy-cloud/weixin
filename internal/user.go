package internal

import "time"

// WechatUser 微信用户表
// 用于存储微信登录的用户信息
type WechatUser struct {
	Id        string    `json:"id" xorm:"pk"`
	OpenID    string    `json:"openid" xorm:"unique index"`
	UnionID   string    `json:"unionid,omitempty" xorm:"index"`
	Name      string    `json:"name,omitempty"`
	Avatar    string    `json:"avatar,omitempty"`
	Cellphone string    `json:"cellphone,omitempty"`
	Disabled  bool      `json:"disabled,omitempty"`
	Created   time.Time `json:"created,omitempty" xorm:"created"`
	Updated   time.Time `json:"updated,omitempty" xorm:"updated"`
}

