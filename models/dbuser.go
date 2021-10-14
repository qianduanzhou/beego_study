package models

type Dbuser struct {
	Id   int    `orm:"pk;auto"`
	Name string `orm:"size(100)"`
}
