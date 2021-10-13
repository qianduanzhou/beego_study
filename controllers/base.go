package controllers

import (
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

type BaseController struct {
	beego.Controller
	o orm.Ormer
}

func (base *BaseController) Prepare() {
	base.o = orm.NewOrm()
}
