package controllers

import (
	"beego_study/models"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

const (
	NormalCode = 200 //正常
	ErrorCode  = 500 //错误
)

type BaseController struct {
	beego.Controller
	o orm.Ormer
}

type CommonResult struct {
	code int
	data interface{}
	msg  string
}

func (base *BaseController) Prepare() {
	base.o = models.GetO()
}

//获取页面参数
func GetPageParam(l *ListController) (int, int) {
	page, _ := l.GetInt("page")
	limit, _ := l.GetInt("limit")
	return page, limit
}

//构造
func newCommonResult(code int, data interface{}, msg string) *CommonResult {
	return &CommonResult{
		code: code,
		data: data,
		msg:  msg,
	}
}

func GetCommonResult(code int, data interface{}, msg string) *CommonResult {
	common := newCommonResult(code, data, msg)
	var result = make(map[string]interface{})
	result["code"] = common.code
	result["data"] = common.data
	result["msg"] = common.msg
	return common
}
