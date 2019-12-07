package main

import "go_project/gongchang/factory"
import _ "go_project/gongchang/cls1"
import _ "go_project/gongchang/cls2"

func main() {
	c1 := factory.Create("Class1")
	c2 := factory.Create("Class2")

	c1.Do()
	c2.Do()
}