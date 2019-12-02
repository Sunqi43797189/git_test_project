package main

import "fmt"

type Voker interface {
	Call(interface{})
}

type FuncType func(interface{}) // 函数定义为类型才可以实现接口

func main() {
	var voker Voker
	funcType := FuncType(func(v interface{}) {
		fmt.Println("from function", v)
	})
	voker = funcType
	voker.Call("hello")
}

// 实现接口
func (f FuncType) Call(p interface{}) {
	f(p)
	fmt.Println("func interface")
}
