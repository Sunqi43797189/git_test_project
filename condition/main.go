package main

import "fmt"

func main(){
	a := 1
	res := ifTest(a)
	fmt.Println(res)
}

func ifTest(a int) bool {
	if a == 1{
		return true
	} else {
		return false
	}
}
