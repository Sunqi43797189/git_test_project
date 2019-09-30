package main

import "fmt"


func channelTestV2() {
	c := make(chan int)

	go func() {

		c <- 1
		c <- 2
		c <- 3
		close(c)
	}()

	for v := range  c {
		fmt.Println(v)
	}
}
