package main

import "container/list"

import "fmt"

func main()  {
	l := list.New() // 创建列表
	fmt.Println(l)

	l.PushBack("fist") //在列表尾部插入
	l.PushFront(67) // 列表头部插入

	fmt.Println(l)

	element := l.PushFront("element") // 头部插入并保存插入的元素
	l.InsertAfter("high", element) // 在元素的后边插入
	l.InsertBefore("noon", element) // 在元素的前边插入

	l.Remove(element) // 删除元素

	// 遍历元素
	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
	
}