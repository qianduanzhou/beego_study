package controllers

import (
	"beego_study/models"
	"fmt"
)

type ListController struct {
	BaseController
}

// @Title 获取列表
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

// @Title 插入测试列表
// @Description 插入列表，提供测试数据
// @router /insertTestList [post]
func (l *ListController) InsertTestList() {
	n, _ := l.GetInt("num")
	list := models.List{}
	for i := 1; i <= n; i++ {
		list.Id = i
		list.Title = fmt.Sprintf("测试标题%d", i)
		l.o.Insert(&list)
	}
	result := GetCommonResult(NormalCode, "插入成功", "success")
	l.Data["json"] = result
	l.ServeJSON()
}
