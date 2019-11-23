package main

import "fmt"
//var (
//	count int32
//	wg    sync.WaitGroup
//)
//
//func main() {
//	wg.Add(2)
//	go incCount()
//	go incCount()
//	wg.Wait()
//	fmt.Println(count)
//
//}
//
//func incCount() {
//	defer wg.Done()
//	for i := 0; i < 1000; i++ {
//		count ++
//	}
//}

type MyStruct struct {
	name string
	age int
}

var ss MyStruct


func main() {
	//c := make(chan bool)
	//m := make(map[string]string)
	//go func() {
	//	m["1"] = "a" // First conflicting access.
	//	c <- true
	//}()
	//m["2"] = "b" // Second conflicting access.
	//<-c
	//for k, v := range m {
	//	fmt.Println(k, v)
	//}
	ss.name = "111"
	ss.age = 11
	fmt.Println(ss)
}