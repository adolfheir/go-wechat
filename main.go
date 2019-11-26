package main

import (
	"fmt"
	"log"
	"syscall"
	"time"

	wechat "github.com/adolfheir/go-wechat/controller/wx"
	"github.com/adolfheir/go-wechat/router"
	"github.com/adolfheir/go-wechat/setting"
	"github.com/chanxuehong/wechat/mp/menu"
	"github.com/fvbock/endless"
)

// @title Go WX API
// @version 1.0
// @description for learn go && wx
// @termsOfService https://github.com/adolfheir/go-wechat
func main() {
	fmt.Println(setting.ServerSetting.Ip)
	r := router.InitRouter()

	//生成微信公众号按钮
	menu.Delete(wechat.WechatClient)

	var button1 menu.Button
	button1.SetAsClickButton("点击测试", "test_click")
	var button2 menu.Button
	button2.SetAsClickButton("点击测试2", "test_click2")

	var buttons []menu.Button
	buttons = append(buttons, button1)
	buttons = append(buttons, button2)

	_menu := new(menu.Menu)
	_menu.Buttons = buttons
	menu.Create(wechat.WechatClient, _menu)
	time.Sleep(time.Duration(2) * time.Second)
	fmt.Println(menu.GetMenuInfo(wechat.WechatClient))

	/**
	启动服务器
	*/
	address := fmt.Sprintf("%s:%s", setting.ServerSetting.Ip, setting.ServerSetting.Port)
	server := endless.NewServer(address, r)
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}
	// 处理服务器错误
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}
