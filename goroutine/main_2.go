package main

import (
	"fmt"
	"sync"
	"runtime"
	"time"
)

func main(){
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println(runtime.NumCPU())
	var wg = sync.WaitGroup{}
	wg.Add(4)
	go func() {
		defer wg.Done()
		for i:= 0; i < 100; i++ {
			fmt.Println("A", i)
			time.Sleep(time.Second * 1)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			fmt.Println("B", i)
			time.Sleep(time.Second * 1)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			fmt.Println("C", i)
			time.Sleep(time.Second * 1)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			fmt.Println("D", i)
			time.Sleep(time.Second * 1)
		}
	}()
	wg.Wait()
}
