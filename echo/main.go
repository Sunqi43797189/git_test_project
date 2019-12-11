package main

import "os"
import "go_project/echo/server"

func main()  {
	exitChan := make(chan int)
	go server.Server("127.0.0.1:12306", exitChan)
	code := <- exitChan
	fmt
	os.Exit(code)
}
