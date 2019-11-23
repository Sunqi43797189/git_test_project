package main

import "fmt"

func do(i interface{}){
	switch v := i.(type) {
	case int:
		fmt.Println("int", v)
	case string:
		fmt.Println("string", v)
	case float64:
		fmt.Println("float", v)
	default:
		fmt.Println("default")
	}
}

func main(){
	do(22)
	do("33")
	do(22.2)
	do(map[string]string{"name":"sunqi"})
}
