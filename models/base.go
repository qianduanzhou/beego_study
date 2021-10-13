package models

import (
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	// set default database
	orm.RegisterDriver("mysql", orm.DRMySQL)                                                                   //使用mysql数据库
	orm.RegisterDataBase("default", "mysql", "root:123456@tcp(127.0.0.1:3306)/go_test?charset=utf8&loc=Local") //注册mysql（default为数据库别名）
	orm.RegisterModel(new(Dbuser))                                                                             //设置model
	orm.RunSyncdb("default", false, true)                                                                      //根据model创建表
}
