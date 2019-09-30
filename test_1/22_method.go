package main

import (
	"fmt"
	"image/color"
	"math"
)

type Point struct{
	X, Y float64
}

type Path []Point

type IntList struct {
	Value int
	Tail *IntList
}

type Color struct {
	Point
	Color color.RGBA
}


//传统方法
func Distance(p, q Point) float64{
	return math.Hypot(q.X - p.X, q.Y - p.Y)
}

//Point 类型的方法
func (p Point)Distance(q Point) float64 {
	return math.Hypot(q.X - p.X, q.Y - p.Y)
}

//func (path Path)Distance() float64 {
//	sum := 0.0
//	for i := range path {
//		if i > 0 {
//			sum += path[i-1].Distance(path[i])
//		}
//	}
//  return sum
//}

func (path *Path)Distance() float64 {
	sum := 0.0
	for i := range *path {
		if i > 0 {
			sum += (*path)[i-1].Distance((*path)[i])
		}
	}
	return sum
}

func (list *IntList) Sum() int {
	if list == nil {
		return 0
	}
	return list.Value + list.Tail.Sum()
}

func methodTest() {
	p := Point{X: 1,Y: 2}
	q := Point{X: 4,Y: 6}
	fmt.Println(Distance(p, q))
	fmt.Println(p.Distance(q))

	red := color.RGBA{255, 0, 0, 255}
	c := Color{p, red}
	fmt.Println(c.Distance(q)) //嵌套结构体也可以调用内部结构体的方法
}