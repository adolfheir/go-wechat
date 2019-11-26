package setting

import (
	"log"

	"gopkg.in/ini.v1"
)

type AppSetting struct {
	JwtSecret string
	WxToken   string
}
type ServerSetting struct {
	Ip   string
	Port string
}
type DatabaseSetting struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

var App = &AppSetting{}
var Server = &ServerSetting{}
var Database = &DatabaseSetting{}
var config *ini.File

func init() {
	var err error
	config, err = ini.Load("config/app.ini")
	if err != nil {
		log.Fatal("Fail to parse 'config/app.ini': %v", err)
	}
	mapTo("app", App)
	mapTo("server", Server)
	mapTo("database", Database)
}

func mapTo(section string, v interface{}) {
	err := config.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo RedisSetting err: %v", err)
	}
}
