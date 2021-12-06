package controllers

import "fmt"

type ListController struct {
	BaseController
}

// @Title get
// @Description 查询列表
// @router /get [get]
func (l *ListController) Get() {
	fmt.Println("ListController")
	page, limit := GetPageParam(l)
	fmt.Println(page, limit)
	data := GetCommonResult(NormalCode, "test", "success")
	l.Data["json"] = *data
	fmt.Println(*data)
	l.ServeJSON()
}
