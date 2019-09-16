package main

import (
	"fmt"
	"os"
	"strings"
)
func echoTest() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = ""
	}
	fmt.Println(s)
}

func echoTestV2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func echoTestV3() {
	fmt.Println(strings.Join(os.Args[1:], ""))
}

func echoTestV4() {
	fmt.Println(os.Args[1:])
}

func echoTestV5Index() {
	for index, arg := range os.Args {
		fmt.Println(index, arg)
	}
}