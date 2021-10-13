// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"beego_study/controllers"
	"fmt"

	web "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

func init() {
	/*基础路由*/
	//基本 GET 路由
	web.Get("/", func(ctx *context.Context) {
		ctx.Output.Body([]byte("hello world"))
	})
	//基本 POST 路由
	web.Post("/alice", func(ctx *context.Context) {
		ctx.Output.Body([]byte("bob"))
	})
	//注册一个可以响应任何 HTTP 的路由
	web.Any("/foo", func(ctx *context.Context) {
		ctx.Output.Body([]byte("bar"))
	})

	//支持自定义的 handler 实现
	// s := rpc.NewServer()
	// s.RegisterCodec(json.NewCodec(), "application/json")
	// s.RegisterService(new(HelloService), "")
	// web.Handler("/rpc", s)

	/*RESTful Controller 路由*/
	//固定路由(固定路由也就是全匹配的路由，如下所示：)
	// web.Router("/", &controllers.MainController{})
	// web.Router("/admin", &admin.UserController{})
	// web.Router("/admin/index", &admin.ArticleController{})
	// web.Router("/admin/addpkg", &admin.AddController{})

	//正则路由
	//为了用户更加方便的路由设置，beego 参考了 sinatra 的路由实现，支持多种方式的路由：
	// web.Router("/api/?:id", &controllers.TestContrller{})
	//自定义
	// web.Router("/test", &controllers.TestContrller{}, "get:Test")
	//自动匹配
	// web.AutoRouter(&controllers.TestContrller{})
	//注解路由
	// web.Include(&controllers.TestContrller{})
	//方法表达式路由
	// web.CtrlGet("/ping", &controllers.TestContrller{}.Ping)
	//命名空间
	ns := web.NewNamespace("/v1",
		web.NSNamespace("/object",
			web.NSInclude(
				&controllers.ObjectController{},
			),
		),
		web.NSNamespace("/user",
			web.NSInclude(
				&controllers.UserController{},
			),
		),
	)
	web.AddNamespace(ns)

	//form参数解析
	web.Router("/parseForm", &controllers.FormControllers{})
	//下载文件
	web.AddNamespace(web.NewNamespace("/file", web.NSInclude(&controllers.FileController{})))

	//session
	web.Router("/getSession", &controllers.TestContrller{}, "get:GetSess")

	//过滤器
	var FilterUser = func(ctx *context.Context) {
		_, ok := ctx.Input.Session("uid").(int)
		fmt.Println(ok)
		if !ok && ctx.Request.RequestURI != "/login" {
			ctx.Redirect(302, "/login")
		}
	}
	web.InsertFilter("/index", web.BeforeRouter, FilterUser)

	//数据库操作
	web.AddNamespace(web.NewNamespace("/dbuser", web.NSInclude(&controllers.DbuserController{})))
}
