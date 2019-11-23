package main

import "fmt"

type Bag struct {
	items []int
}

type Point struct {
	X int
	Y int
}

func main() {
	//var bag = &Bag{items: []int{1,2,3}}
	//fmt.Println(bag)
	//insertBag(bag, 4)
	//fmt.Println(bag.items)
	//bag.insertBagV2(5)
	//fmt.Println(bag.items)

	var p = new(Point)
	p.X = 1
	p.Y = 2
	fmt.Println(p)
	q := p.setValue(3,4)
	fmt.Println(p)
	fmt.Println(q)
}

func insertBag(b *Bag, itemId int){
	b.items = append(b.items, itemId)
}

func (b *Bag)insertBagV2(itemId int){
	// 指针接收器
	b.items = append(b.items, itemId)
}

func (p Point)setValue(x,y int) *Point{
	// 非指针接收器 对 接收器进行一个拷贝
	p.X = x
	p.Y = y
	return &p
}