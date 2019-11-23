package main

import "fmt"

func main(){
	var x interface{} = 123

	n, ok := x.(int)
	fmt.Println(n, ok)

	n = x.(int)
	fmt.Println(n)

	a, ok := x.(float64)
	fmt.Println(a, ok)

	a = x.(float64) // 产生异常 interface conversion: interface {} is int, not float64
}
