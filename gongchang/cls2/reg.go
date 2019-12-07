package cls2

import "fmt"

import "go_project/gongchang/factory"

type Class2 struct {
	name string
}

func (c *Class2) Do() {
	fmt.Println(c.name)
}

func init() {
	fmt.Println("cls2 register")
	factory.Register("Class2", func() factory.Class {
		class2 := Class2{name: "liyanjin"}
		return &class2
	})
}
