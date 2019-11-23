package main

import "fmt"

func main(){
	var a [5]int // 初始化 0值 [0 0 0 0 0]
	//fmt.Println(a)
	//
	//b := [5]*int{1:new(int), 3: new(int)}
	//*b[0] = 1
	//fmt.Print(b)

	funcA(&a)
	fmt.Println(a)
}

func funcA(a *[5]int){
	a[0] = 1
	fmt.Println(a)
}
