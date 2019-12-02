package main

import "fmt"

func main() {
	defer fmt.Println("1111")
	defer func()  {
		err := recover()
		fmt.Println("err is ", err)
	}()
	panic("panic")
}