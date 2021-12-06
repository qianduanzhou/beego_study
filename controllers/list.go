package controllers

import (
	"beego_study/models"
	"fmt"
)

type ListController struct {
	BaseController
}

// @Title get
// @Description 查询列表
// @router /get [get]
func (l *ListController) Get() {
	page, limit := GetPageParam(l)
	fmt.Println(page, limit)
	list := []models.List{}
	if page != 0 && limit != 0 {
		qs := l.o.QueryTable("list")
		qs.Limit(limit).Offset((page - 1) * limit).All(&list)
	} else {
		l.o.QueryTable("list").All(&list)
	}
	var count int = 0
	l.o.Raw("SELECT COUNT(*) FROM list").QueryRow(&count)
	data := map[string]interface{}{
		"record":     list,
		"totalCount": count,
	}
	result := GetCommonResult(NormalCode, data, "success")
	l.Data["json"] = result
	l.ServeJSON()
}
