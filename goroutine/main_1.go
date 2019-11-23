package main

import (
	"time"
	"fmt"
)

//func running(){
//	times := 0
//
//	for{
//		times ++
//		fmt.Println("time ", times)
//		time.Sleep(time.Second)
//	}
//}

//func main(){
//	go running()
//	go func() {
//		times := 0
//		for {
//			fmt.Println("time * 2", times)
//			times = times + 2
//			time.Sleep(time.Second * 2)
//		}
//	}()
//	var input string
//	res, err := fmt.Scanln(&input)
//	if err != nil{
//		fmt.Println("error")
//	}
//	fmt.Println(res)
//}


//func main(){
//	ch := make(chan int)
//	go func() {
//		fmt.Println("start ...")
//		ch <- 0
//		fmt.Println("end ...")
//	}()
//
//	fmt.Println("waiting ...")
//	res :=<- ch
//	fmt.Println(res)
//	fmt.Println("all done...")
//}


//func printer(c chan int){
//	for{
//		data := <-c
//		if data == 0 {
//			break
//		}
//		fmt.Println(data)
//	}
//	c <- 0
//}
//
//func main(){
//	c := make(chan int)
//	go printer(c)
//
//	for i := 1; i <= 10; i ++ {
//		c <- i
//	}
//	c <- 0
//	res :=<- c
//	fmt.Println(res)
//}


//func main(){
//	timeout := make(chan int, 1)
//	ch := make(chan int)
//	fmt.Println(timeout)
//
//	go func() {
//		time.Sleep(time.Second)
//		timeout <- 2
//	}()
//
//	select {
//		case <-ch:
//			fmt.Println("read data")
//		case <-timeout:
//			fmt.Println("timeout")
//	}
//}

//func main(){
	//ch := make(chan int,1)
	//go func() {
	//	for{
	//		fmt.Println(<- ch)
	//	}
	//}()
	//for {
	//	select {
	//	    case ch <- 0:
	//	    	time.Sleep(time.Second)
	//		case ch <- 1:
	//			time.Sleep(time.Second)
	//	}
	//}
//}

//func say(s string) {
//	for i := 0; i < 5; i++ {
//		time.Sleep(time.Second * 5)
//		fmt.Println(s)
//	}
//}

//func millionGroutine(){
//	for  i := 0; i <= 1000000;i++{
//		go func(i int) {
//			fmt.Println("go func ", i)
//			time.Sleep(time.Second * 2)
//		}(i)
//	}
//}
//
//func main() {
//	millionGroutine()
//}

//func main() {
//	userCount := math.MaxInt64
//	for i := 0; i < userCount; i++ {
//		go func(i int) {
//			// 做一些各种各样的业务逻辑处理
//			fmt.Printf("go func: %d\n", i)
//			time.Sleep(time.Second)
//		}(i)
//	}
//}

var sema = gsema.NewSemaphore(3)

func main(){
	userCount := 10
	for i := 0; i < userCount; i++ {
		go Read(i)
	}
}

func Read(i int){
	defer sema.Done()
	sema.Add(1)

	fmt.Println("go func %d\n, time: %d\n", i, time.Now().Unix())
	time.Sleep(time.Second * 2)
}
