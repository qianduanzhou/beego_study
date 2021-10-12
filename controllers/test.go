package controllers

import (
	"fmt"

	beego "github.com/beego/beego/v2/server/web"
)

type TestContrller struct {
	beego.Controller
}

//获取参数
func (t *TestContrller) Get() {
	param := t.Ctx.Input.Param(":id")
	var str string = fmt.Sprintf("获取参数：%v", param)
	t.Ctx.WriteString(str)
}

//自定义
func (t *TestContrller) Test() {
	t.Ctx.WriteString("test")
}

//注解路由
// @router /zhujie [get]
func (t *TestContrller) Zhujie() {
	t.Ctx.WriteString("注解路由")
}

//方法表达式路由
func (t *TestContrller) Ping() {
	t.Data["json"] = "pong"
	t.ServeJSON()
}

//session
func (t *TestContrller) GetSess() {
	fmt.Println("GetSess")
	v := t.GetSession("asta")
	if v == nil {
		t.SetSession("asta", int(1))
		t.Ctx.WriteString("SetSess")
	} else {
		t.SetSession("asta", v.(int)+1)
		t.Ctx.WriteString("GetSess")
	}
}
