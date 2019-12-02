package main

import "fmt"

func main() {
	var i int
	// for {
	// 	f := AddValue(i)
	// 	res := f()
	// 	fmt.Println(res)
	// 	time.Sleep(time.Second * 2)
	// 	i++
	// }
	f := AddValue(i)
	fmt.Println(f())
	fmt.Println(f())

}

func AddValue(value int ) func() int {
	return func() int {
		value++
		return value
	}
}

