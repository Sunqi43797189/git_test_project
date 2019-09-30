package main

import "fmt"

type Person struct {
	name string
	age int
}

func (p Person)sayInfo() {
	fmt.Println(p.name)
	fmt.Println(p.age)
}

func (per *Person)setAge(age int) {
	fmt.Println(per.age)
	per.age = age
	fmt.Println(per.age)
}

func funcTest() {
	p := Person{"sunqi", 12}
	p.sayInfo()
	p.setAge(15)
	fmt.Println(p)
}