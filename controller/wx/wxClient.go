package wechat

import (
	"github.com/adolfheir/go-wechat/setting"
	"github.com/chanxuehong/wechat/mp/core"
)

var (
	AccessTokenServer core.AccessTokenServer
	WechatClient      *core.Client
)

func init() {
	var (
		wxAppId     = setting.AppSetting.WxAppID
		wxAppSecret = setting.AppSetting.WXAppSecret
	)
	AccessTokenServer = core.NewDefaultAccessTokenServer(wxAppId, wxAppSecret, nil)
	WechatClient = core.NewClient(AccessTokenServer, nil)
}
