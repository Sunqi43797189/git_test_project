package session

import "net"

import "fmt"

import "bufio"

import "strings"
import "go_project/echo/telnet"

func HandleSession(conn net.Conn, exitChan chan int) {
	defer conn.Close()
	fmt.Println("session started...")
	reader := bufio.NewReader(conn)
	for {
		str, err := reader.ReadString('\n')
		if err == nil {
			str = strings.TrimSpace(str)
			if !telnet.ProcessTelnetCommand(str, exitChan){
				return
			}
			conn.Write([]byte(str + "\r\n"))
		} else {
			fmt.Println("session closed ...")
			return
		}
	}	
}
