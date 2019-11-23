package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	channel := make(chan string)
	go producer("dog", channel)
	go producer("cat", channel)
	consumer(channel)
}

func producer(header string, channel chan string){
	for {
		channel <- fmt.Sprintf("%s, %v", header, rand.Int31())
		time.Sleep(time.Second)
	}
}

func consumer(channel chan string){
	for{
		str := <-channel
		fmt.Println(str)
	}
}