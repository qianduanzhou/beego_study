package models

type Form struct {
	name string `form:"username"`
	age  int    `form:"age"`
}

func GetForm(name string, age int) (f *Form) {
	f = &Form{
		name,
		age,
	}
	return
}
