package main

import (
	"fmt"
	"math"
)

func iotaTest()  {
	const (
		a = iota
		b
		c
		d = "ha"
		e
		f = 100
		g
	)

	const (
		_ = iota // 重置iota
		i
	)
    // 常量声明时 ，如果未定义初始值 则会使用上一行常量的右值
	println(a,b,c,d,e,f,g,i)
}

func simpleIota() {
	type Weekday int

	const (
		Sunday Weekday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)

	const (
		_ = 1 << (10 * iota)
		KB
		MB
		GB
		TB
		PB
		EB
		ZB
		YB
	)

	var x float32 = math.Pi
	fmt.Println(x)
	var y float64 = math.Pi
	fmt.Println(y)
	var z complex128 = math.Pi
	fmt.Println(z)
}
