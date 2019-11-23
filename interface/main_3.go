package main

import "fmt"

type animal interface {
	printInfo()
}

type cat int
type dog int

func (c cat) printInfo(){
	fmt.Println("cat")
}

func (d *dog)printInfo() {
	fmt.Println("dog")
}
func main() {
	var c cat
	var d dog
	invoke(c)
	invoke(&c)
	invoke(&d)
	//invoke(d) // 报错
}

func invoke(a animal) {
	a.printInfo()
}
