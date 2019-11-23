package main

func closureTest() {
	// var str string = "hello World"
	// var f = func() {
	// 	str = "hello" + str
	// }

	// sunqi := f()
	// tom := f()
	// jerry := f()
}

func Adder() func(int) int {
	var x int
	return func(delta int) int {
		x += delta
		return x
	}
}
