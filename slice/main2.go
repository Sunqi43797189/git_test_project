package main

import (
	"fmt"
)

// func main() {
// 	z := []int{1,2,3,4}
// 	a := z
// 	b := []int{0,0,0,0}

// 	fmt.Println(z, a, b)

// 	num := copy(b, a)
	
// 	fmt.Println(num)
// 	fmt.Println(z, a, b)
	
// 	z[0] = 100
// 	// a[0] = 10
// 	fmt.Println(z, a, b)

// }

func main()  {
	a := []int{1, 2, 3, 4, 5, 6}
	seq := append(a[:1], a[2:]...)
	fmt.Println(seq)
}