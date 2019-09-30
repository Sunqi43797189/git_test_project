package main

import "fmt"

func chengFaBiao(){
	for y := 0; y <= 9; y++ {
		for x:= 0; x <= y ;x++ {
			fmt.Printf(" %d * %d = %d ", x, y, x * y)
		}
		fmt.Printf("\n")
	}
}