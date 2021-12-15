package models

import (
	"fmt"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/config"
	_ "github.com/go-sql-driver/mysql"
)

var O orm.Ormer

func init() {
	mysqladmin, _ := config.String("mysqladmin")
	mysqlPwd, _ := config.String("mysqlPwd")
	mysqlDB, _ := config.String("mysqlDB")
	mysqlUrl, _ := config.String("mysqlUrl")
	mysqlCharset, _ := config.String("mysqlCharset")
	fmt.Println(mysqladmin, mysqlPwd, mysqlDB, mysqlUrl, mysqlCharset)
	orm.RegisterDriver("mysql", orm.DRMySQL)                                                                                                    //使用mysql数据库
	orm.RegisterDataBase("default", "mysql", fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=%v", mysqladmin, mysqlPwd, mysqlUrl, mysqlDB, mysqlCharset)) //注册mysql（default为数据库别名）
	orm.RegisterModel(new(Dbuser))                                                                                                              //设置model
	orm.RegisterModel(new(List))                                                                                                                //设置model
	orm.RunSyncdb("default", false, true)                                                                                                       //根据model创建表
	O = orm.NewOrm()
}

func GetO() orm.Ormer {
	return O
}
