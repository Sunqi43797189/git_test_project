package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

func main() {
	s := "hello 你好"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))

	// for _, value := range(s) {
	// 	fmt.Println(string(value))
	// }

	// for i := 0; i < len(s); i++ {
	// 	fmt.Println(string(s[i]))
	// }
	stringToInt()
}

func stringToInt() {
	str := "1111"
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(num)
	}
}
