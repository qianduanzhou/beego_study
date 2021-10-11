package models

type Form struct {
	Name string `form:"username"`
	Age  int    `form:"age"`
}

func GetForm(name string, age int) (f *Form) {
	f = &Form{
		name,
		age,
	}
	return
}
