package main

import (
	_ "beego_study/routers"
	"fmt"

	"github.com/astaxie/beego/session"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

var globalSessions *session.Manager

func initSession() {
	sessionConfig := &session.ManagerConfig{
		CookieName:      "uid",
		EnableSetCookie: true,
		Gclifetime:      3600,
		Maxlifetime:     3600,
		Secure:          false,
		CookieLifeTime:  3600,
		// ProviderConfig:  "root:123456@tcp(127.0.0.1:3306)/go_test", //mysql数据库模式
		ProviderConfig: "./tmp", //file模式，session文件存放的地方
	}
	var err error
	globalSessions, err = session.NewManager("file", sessionConfig)
	fmt.Println("globalSessions", globalSessions)
	if err != nil {
		fmt.Println("err", err)
	}
	go globalSessions.GC()
}
func init() {
	// initSession()
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.Run()
}
