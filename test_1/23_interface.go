package main

import "fmt"

type USB interface{
	start()
	end()
}

type Computer struct{
	name string
}

func (c Computer)start(){
	fmt.Println(c.name, "start")
}

func (c Computer)end(){
	fmt.Println(c.name, "end")
}

type Phone struct {
	name string
}

func (ph Phone)start(){
	fmt.Println(ph.name, "start")
}

func (ph Phone)end() {
	fmt.Println(ph.name, "end")
}
func interfaceTest() {
	var in USB
	var cm Computer = Computer{"samsung"}
	Option(cm)
	var ph Phone = Phone{"apple"}
	Option(ph)

	//接口保存变量 取值方法一
	in = cm
	value, ok := in.(Computer)
	fmt.Println(value)
	fmt.Println(ok)
	// 接口保存变量 取值方法二
}

func Option(in USB){
	in.start()
	in.end()
}

func nilInterface() {
	var arr [4]interface{}
	arr[0] = 1
	arr[1] = "sss"
	arr[2] = []int{1,2,3}
	arr[3] = false

	fmt.Println(arr)
}

