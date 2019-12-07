package cls1

import "fmt"

import "go_project/gongchang/factory"

type Class1 struct {
	name string
}

func (c *Class1) Do() {
	fmt.Println(c.name)
}

func init() {
	fmt.Println("cls1 register")
	factory.Register("Class1", func() factory.Class {
		class1 := Class1{name: "sunqi"}
		return &class1
	})
}
