package controllers

import (
	"fmt"
	"log"

	"github.com/beego/beego/v2/server/web"
)

type FileController struct {
	web.Controller
}

// @Title downFile
// @Description 下载文件
// @router /downFile [post]
func (f *FileController) Post() {
	fmt.Println("downFile")
	f2, h, err := f.GetFile("uploadname")
	if err != nil {
		log.Fatal("getfile err ", err)
	}
	defer f2.Close()
	f.SaveToFile("uploadname", "static/upload/"+h.Filename) // 保存位置在 static/upload, 没有文件夹要先创建
}
