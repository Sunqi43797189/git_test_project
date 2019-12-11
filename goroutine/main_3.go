package main

import "time"
import "fmt"

func main () {
	ch := make(chan int)

	// var chSendOnly chan<- int = ch
	// var chRecvOnli <-chan int = ch

	// onlyReadChan := make(<-chan int)
	fmt.Println("start")
	time.AfterFunc(time.Second * 2 , func(){
		fmt.Println(111)
		ch <- 0
	})
	<-ch

	fmt.Println("计时器")
	ticker := time.NewTicker(time.Second * 1)
	stopper := time.NewTimer(time.Second * 5)
	var i int
	for {
		select {
		case <- stopper.C:
			fmt.Println("stop")
			goto StopHere
		case <-ticker.C:
			i++
			fmt.Println("tick", i)
		}
	}
StopHere:
    fmt.Println("done")
}
