package controllers

import (
	"beego_study/models"
	"fmt"

	"github.com/astaxie/beego/validation"
	"github.com/beego/beego/v2/server/web"
)

type FormControllers struct {
	web.Controller
}

//解析自动赋值，需要post且为content-type为application/x-www-form-urlencoded
func (f *FormControllers) Post() {
	form := models.Form{}
	if err := f.ParseForm(&form); err != nil {
		//handle error
		fmt.Println(err)
	}
	f.Data["json"] = form
	f.ServeJSON()

	//绑定age, f.Ctx.Input.Bind方式
	var age int
	f.Ctx.Input.Bind(&age, "age")
	fmt.Println(age)

	//表单验证
	valid := validation.Validation{}
	b, err := valid.Valid(&form)
	if err != nil {
		// handle error
	}
	if !b {
		for _, err := range valid.Errors {
			fmt.Println(err.Key, err.Message)
		}
	}
	fmt.Println(form)

}
