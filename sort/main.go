package main

import (
	"fmt"
	"sort"
)

type MyStringList []string

func (m MyStringList)Len() int {
	return len(m)
}

func (m MyStringList)Less(i, j int) bool {
	return m[i] < m[j]
}

func (m MyStringList)Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func main(){
	names := MyStringList{
		"3. Triple Kill",
		"5. Penta Kill",
		"2. Double Kill",
		"4. Quadra Kill",
		"1. First Blood",
	}
	fmt.Println(names)

	sort.Sort(names)

	fmt.Println(names)

	stringSlice := sort.StringSlice{
		"3. Triple Kill",
		"5. Penta Kill",
		"2. Double Kill",
		"4. Quadra Kill",
		"1. First Blood",
	}
	fmt.Println(stringSlice)
	sort.Sort(stringSlice)
	fmt.Println(stringSlice)
}