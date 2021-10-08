package controllers

import (
	"beego_study/models"
	"fmt"

	"github.com/beego/beego/v2/server/web"
)

type FormControllers struct {
	web.Controller
}

func (f *FormControllers) Post() {
	var form models.Form
	if err := f.ParseForm(&form); err != nil {
		//handle error
		fmt.Println(err)
	}
	f.Data["json"] = form
	f.ServeJSON()
}
