package main

import (
	"fmt"
	"strings"
)

func stringTest()  {
	s := "abcdefgabcdefg"
	b := []byte(s)
	fmt.Println(b)
	s2 := string(b)
	fmt.Println(s2)

	var s_array = []string{"s", "b", "s", "b"}

	s_contain := strings.Contains(s, "a")
	s_count := strings.Count(s, "a")
	s_fields := strings.Fields(s)
	s_hasprefix := strings.HasPrefix(s, "abc")
	s_index := strings.Index(s, "a")
	s_join := strings.Join(s_array, "*")
	fmt.Println(s_contain)
	fmt.Println(s_count)
	fmt.Println(s_fields)
	fmt.Println(s_hasprefix)
	fmt.Println(s_index)
	fmt.Println(s_join)
}