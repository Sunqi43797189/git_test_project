package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func Server(address string, exitChan chan int){
	l, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("error, exit")
		exitChan <- 1
	}
	fmt.Println("listen: " + address)
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
        fmt.Println("connecting ", address)
		go HandleSession(conn, exitChan)
	}
}

func HandleSession(conn net.Conn, exitChan chan int) {
	fmt.Println("session start...")
	reader := bufio.NewReader(conn)
	for {
		str, err := reader.ReadString('\n')
		if err != nil{
			fmt.Println("session close...")
			conn.Close()
			break
		} else {
			str = strings.TrimSpace(str)
			if !processTelnetCommand(str, exitChan){
				conn.Close()
				break
			}
			conn.Write([]byte(str + "\r\n"))
		}
	}
	fmt.Println("shut down ")
}

func processTelnetCommand(str string, exitChan chan int) bool {
	if strings.HasPrefix(str, "@close"){
		fmt.Println("session close ...")
		return false
	} else if strings.HasPrefix(str, "@shutdown"){
		fmt.Println("server shutdown ...")
		exitChan <- 0
		return false
	}
	fmt.Println(str)
	return true
}


func main(){
	exitChan := make(chan int)
	go Server("127.0.0.1:65535",exitChan)
	code := <- exitChan
	os.Exit(code)
}
