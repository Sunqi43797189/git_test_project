package main

import (
	"fmt"
)

func sliceTest() {
	months := [...]string{1: "January", 2: "February", 3: "March", 4: "April", 5: "May", 6: "June", 7: "July", 8: "August", 9: "September", 10: "October", 11: "November", 12: "December"}
	fmt.Println(months[0])

	fmt.Println(cap(months))
	fmt.Println(len(months))

	q2 := months[4:7]
	summer := months[6:9] //切片共享之前的底层数组，因此切片操作对应常量时间复杂度
	fmt.Println(q2)
	fmt.Println(summer)

	for _, q := range q2 {
		for _, s := range summer {
			if s == q {
				fmt.Println(s)
			}
		}
	}
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = j+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func reverseTest() {
	a := [...]int{0, 1, 2, 3, 4, 5}
	s := []int{0, 1, 2, 3, 4, 5}

	fmt.Printf("%T\t%T", a, s)
	//reverse(a[:2])
	//fmt.Println(a)
	//reverse(a[2:])
	//fmt.Println(a)
	//reverse(a)
	//fmt.Println(a)
}

func equalSlice(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func sliceNil() {
	var s []int
	fmt.Println(s == nil)
	fmt.Println(len(s))

	s = nil
	fmt.Println(s == nil)
	fmt.Println(len(s))

	s = []int(nil)
	fmt.Println(s == nil)
	fmt.Println(len(s))

	s = []int{}
	fmt.Println(s == nil)
	fmt.Println(len(s))
}

func makeSlice() {
	a := make([]string, 10) // 长度制定为10
	fmt.Println(a[0])

	b := make([]string, 20)[:10] //容量制定为20，但是去前10个分片，剩余容量留给未来增长
	fmt.Println(b[0])

	fmt.Println(equal(a, b))
}

func appendTest() {
	var runes []rune
	for _, r := range "hello, 世界" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes)
	fmt.Println([]rune("hello, 世界"))

	var x []int
	x = append(x, 1)
	x = append(x, 2, 3)
	x = append(x, 4, 5, 6)
	x = append(x, x...)
	fmt.Println(x)
}

func sliceMemory(strings []string) []string {
	// 原地修改slice内容，返回切片到i的slice
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func sliceMemoryTest() {
	// 原数据被覆盖
	data := []string{"onr", "", "three"}
	fmt.Println(sliceMemory(data))
	fmt.Println(data)
}

func removeV1(slice []int, i int) []int {
	//删除i位置的元素， 保持slice顺序
	copy(slice[i:], slice[i+1:])
	return slice[:len(slice)-1]
}

func removeV2(slice []int, i int) []int {
	// 删除i位置的元素，不保留顺序
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}
