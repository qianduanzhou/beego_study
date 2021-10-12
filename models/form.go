package models

import (
	"strings"

	"github.com/astaxie/beego/validation"
)

type Form struct {
	Name string `form:"username" valid:"Required;Match(/^Bee.*/)"`
	Age  int    `form:"age"`
}

func GetForm(name string, age int) (f *Form) {
	f = &Form{
		name,
		age,
	}
	return
}

//表单验证
func (f *Form) Valid(v *validation.Validation) {
	if strings.Contains(f.Name, "admin") {
		// 通过 SetError 设置 Name 的错误信息，HasErrors 将会返回 true
		v.SetError("Name", "名称里不能含有 admin")
	}
}
