package main

import (
	"fmt"
	"go_project/package_1/package_2"
)

func main() {
	res := package_2.RedisClient()
	fmt.Println(res)
}