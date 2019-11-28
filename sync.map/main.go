package main

import (
	"sync"
	"fmt"
)

func main()  {
	var scene sync.Map
	scene.Store("greece", 97) // 存储
	scene.Store("lodon", 100)
	scene.Store("egypt", 200)

	fmt.Println(scene.Load("lodon")) // 获取

	scene.Delete("lodon") // 删除

	//遍历，匿名函数返回true迭代遍历，false终止遍历
	scene.Range(func(k, v interface{}) bool {
		fmt.Println("iterate:", k, v)
		return false
	})
}