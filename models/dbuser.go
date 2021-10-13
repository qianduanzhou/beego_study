package models

type Dbuser struct {
	Id   int
	Name string `orm:"size(100)"`
}
