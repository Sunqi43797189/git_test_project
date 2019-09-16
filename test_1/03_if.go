package main

import "fmt"

func ifTest()  {
	var a int = 10

	if a < 20 {
		fmt.Printf("a < 20\n")
	}else {
		fmt.Printf("a > 20\n")
	}
	fmt.Printf("a 的值为：%d\n", a)

	if num := 10; num % 2 == 0 {
		fmt.Printf("num is even")
	} else {
		fmt.Printf("num is odd")
	}
}
