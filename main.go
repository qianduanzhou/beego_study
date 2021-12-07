package main

import (
	_ "beego_study/models"
	_ "beego_study/routers"

	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/cors"
)

// var globalSessions *session.Manager

// //初始化session
// func initSession() {
// 	sessionConfig := &session.ManagerConfig{
// 		CookieName:      "uid",
// 		EnableSetCookie: true,
// 		Gclifetime:      3600,
// 		Maxlifetime:     3600,
// 		Secure:          false,
// 		CookieLifeTime:  3600,
// 		// ProviderConfig:  "root:123456@tcp(127.0.0.1:3306)/go_test", //mysql数据库模式
// 		ProviderConfig: "./tmp", //file模式，session文件存放的地方
// 	}
// 	var err error
// 	globalSessions, err = session.NewManager("file", sessionConfig)
// 	fmt.Println("globalSessions", globalSessions)
// 	if err != nil {
// 		fmt.Println("err", err)
// 	}
// 	go globalSessions.GC()
// }

//增加过滤器，解决跨域
func initFilter() {
	web.InsertFilter("*", web.BeforeRouter, cors.Allow(&cors.Options{
		// 允许访问所有源
		AllowAllOrigins: true,
		// 可选参数"GET", "POST", "PUT", "DELETE", "OPTIONS" (*为所有)
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		// 指的是允许的Header的种类
		AllowHeaders: []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		// 公开的HTTP标头列表
		ExposeHeaders: []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		// 如果设置，则允许共享身份验证凭据，例如cookie
		AllowCredentials: true,
	}))
}

func init() {
	// initSession()
	initFilter()
}

func main() {
	if web.BConfig.RunMode == "dev" {
		web.BConfig.WebConfig.DirectoryIndex = true
		web.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	web.Run()
}
