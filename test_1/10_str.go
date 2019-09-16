package main

import (
	"fmt"
	)


func strTest() {
	s := "hello 世界"
	//fmt.Println(len(s))
	//fmt.Println(utf8.RuneCountInString(s))

	//for i := 0; i < len(s); {
	//	r, size := utf8.DecodeRuneInString(s[i:])
	//	fmt.Printf("%d\t%c\n", i, r)
	//	i += size
	//}

	//for i, r := range "hello 世界" {
	//	fmt.Printf("%d\t%q\t%d\n", i, r, r)
	//}
	//
	//n := 0
	//for range s {
	//	n++
	//}
	//fmt.Println(n)
	r := []rune(s) //编码数组
	fmt.Println(r)
	fmt.Println(string(r)) // 	编码转换字符串
}
