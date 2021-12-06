package controllers

import (
	"beego_study/models"
	"errors"
	"fmt"

	"github.com/beego/beego/v2/client/orm"
)

type DbuserController struct {
	BaseController
}

// @Title search
// @Description 查询
// @Param	id		path 	string	true		"用户的id"
// @router /search [get]
func (dbuser *DbuserController) Search() {
	id, _ := dbuser.GetInt("id")
	user := models.Dbuser{
		Id: id,
	}
	err := dbuser.o.Read(&user)
	if errors.Is(err, orm.ErrNoRows) {
		dbuser.Ctx.WriteString("查询不到")
	} else if errors.Is(err, orm.ErrMissPK) {
		dbuser.Ctx.WriteString("找不到主键")
	} else {
		dbuser.Data["json"] = user
		fmt.Println(user)
		dbuser.ServeJSON()
	}
}

// @Title search2
// @Description 查询 QueryTable
// @Param	id		path 	string	true		"用户的id"
// @router /search2 [get]
func (d *DbuserController) Search2() {
	id, _ := d.GetInt("id")
	user := models.Dbuser{}
	qs := d.o.QueryTable(&user)
	qs.Filter("id", id).One(&user)
	d.Data["json"] = user
	d.ServeJSON()
}

// @Title search3
// @Description 查询 原生sql
// @Param	id		path 	string	true		"用户的id"
// @router /search3 [get]
func (d *DbuserController) Search3() {
	id, _ := d.GetInt("id")
	name := d.GetString("name")
	user := models.Dbuser{}
	users := []models.Dbuser{}
	if err := d.o.Raw("SELECT * FROM dbuser WHERE id = ?", id).QueryRow(&user); err != nil {
		fmt.Println(err)
		d.Ctx.WriteString("找不到")
	} else {
		d.Data["json"] = user
		d.ServeJSON()
	}
	if num, err := d.o.Raw("SELECT * FROM dbuser WHERE name = ?", name).QueryRows(&users); err != nil {
		fmt.Println(err)
		d.Ctx.WriteString("找不到")
	} else {
		fmt.Println(users, num)
		d.Data["json"] = users
		d.ServeJSON()
	}
}

// @Title insert
// @Description 插入
// @router /insert [post]
func (dbuser *DbuserController) Insert() {
	var user = models.Dbuser{}
	user.Name = dbuser.GetString("name")
	id, err := dbuser.o.Insert(&user)
	if err == nil {
		fmt.Println(id)
	} else {
		fmt.Println(err)
	}
}

// @Title update
// @Description 更新
// @router /update [post]
func (dbuser *DbuserController) Update() {
	fmt.Println(dbuser.GetInt("id"))
	var user models.Dbuser
	dbuser.Ctx.Input.Bind(&user.Id, "id")
	if err := dbuser.o.Read(&user); err == nil {
		user.Name = dbuser.GetString("name")
		if num, err := dbuser.o.Update(&user, "Name"); err == nil {
			fmt.Println(num)
		} else {
			fmt.Println("err", err)
		}
	} else {
		fmt.Println("err", err)
	}
}

// @Title delete
// @Description 删除
// @router /delete [post]
func (dbuser *DbuserController) Delete() {
	fmt.Println(dbuser.GetInt("id"))
	var user models.Dbuser
	dbuser.Ctx.Input.Bind(&user.Id, "id")
	if num, err := dbuser.o.Delete(&user); err == nil {
		fmt.Println(num)
	}
}
