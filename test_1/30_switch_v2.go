package main

import "fmt"

func switchTestV2() {
	var a int = 60

	switch  {
	case a < 50:
		fmt.Println("a < 50")
	case a < 60:
		fmt.Println("a < 60")
	case a < 70:
		fmt.Println("a < 70")
	default:
		fmt.Println("default")
	}
}
