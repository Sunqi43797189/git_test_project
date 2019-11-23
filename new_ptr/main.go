package main

import (
	"fmt"
)

func main() {
	str := new(string)
	fmt.Println(*str) // 零值 字符串 空
	*str = "ninja"
	fmt.Println(str, *str)
}