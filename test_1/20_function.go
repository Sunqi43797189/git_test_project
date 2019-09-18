package main

import "fmt"

func squares() func()  {
	var x int
	return func() {
		x++
		fmt.Println(x * x)
	}
}