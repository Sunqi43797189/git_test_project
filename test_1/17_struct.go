package main

import "fmt"

type Address struct {
	hostname string
	port int
}

func structTest() {
	address1 := Address{hostname: "192.168.1.1", port: 9090}
	hits := make(map[Address]int)
	fmt.Println(address1)
	hits[address1]++
	fmt.Println(address1.hostname)
	p := &address1
	fmt.Println(p.port)
	fmt.Println(structUpdate(p).port)
}

func structUpdate(p *Address) *Address {
	p.port++
	return p
}