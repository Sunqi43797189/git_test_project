package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age int
}

func reflectTypeTest(value interface{}) interface{} {
	res := reflect.TypeOf(value)
	return res
}

func reflectValueTest(value interface{}) interface{} {
	res := reflect.ValueOf(value)
	return res
}

func main(){
	//res := reflectTypeTest("str")
	//fmt.Print(res)
	//
	//value := reflectValueTest("str")
	//fmt.Print(value)
	p1 := Person{Name:"sunqi", Age: 30}
	t1 := reflect.TypeOf(p1)
	fmt.Println(t1)
	fmt.Println(t1.Name())
	fmt.Println(t1.Kind())
}