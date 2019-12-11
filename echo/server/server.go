package server

import "net"

import "fmt"
import "go_project/echo/session"

func Server(address string , exitChan chan int)  {
	l, err := net.Listen("tcp", address)
	fmt.Println("start listen 12306")
	if err != nil {
		fmt.Println("链接失败")
		exitChan <- 1
	}

	defer l.Close()
	for {
		conn, err := l.Accept()
		fmt.Println(conn.RemoteAddr())
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		go session.HandleSession(conn, exitChan)
	}
}