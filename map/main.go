package main

import "fmt"

func main()  {
	dict := map[string]int{
		"name": 1,
	}
	fmt.Printf("%p, %v\n", dict, dict)
	map1(dict)
	map2(dict)
}

func map1(m map[string]int)  {
	fmt.Printf("%p, %v\n", m, m)
	m["name"] = 2
	fmt.Printf("%p, %v\n", m, m)
}

func map2(m map[string]int) {
	fmt.Printf("%p, %+v\n", m, m)
}