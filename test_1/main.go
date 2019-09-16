package main

import (
	//"fmt"
	"fmt"
	"strconv"
)
func main() {
	//constTest()
	//iotaTest()
	//ifTest()
	//switch_test()
	//switch_fallthrough()
	//echoTest()
	//echoTestV2()
	//echoTestV4()
	//echoTestV5Index()
	//bufioTest()
	//bufioTestV2()
	//easy_variable_test()
	//strTest()
	//s := baseName("a/b/c.go")
	//fmt.Println(s)
	//s := baseNameV2("a/b/c.go")
	//fmt.Println(s)
	//s := comma("123456789123456789")
	//fmt.Println(s)
	//stringTest()
	//fmt.Println(intsToString([]int{1,2,3}))
	//fmt.Println(commaV2("123456789123456789"))
	//simpleTest()
	//simpleIota()
	//simpleInt()
	//arrayTest()
	//sliceTest()
	//reverseTest()
	//sliceNil()
	//makeSlice()
	appendTest()
}

func simpleTest() {
	x, err := strconv.Atoi("123fa")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(x)
}

func simpleInt() {
	var f float64 = 212
	fmt.Println(f)
	fmt.Println((f-32) * 5 /9)
	fmt.Println(5 / 9 * (f-32))
	fmt.Println(5.0 / 9 * (f-32))
	fmt.Printf("%T", f)
}


