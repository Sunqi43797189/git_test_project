package telnet

import "strings"

import "fmt"

func ProcessTelnetCommand(str string, exitChan chan int) bool {
	if strings.HasPrefix(str, "@close"){
		fmt.Println("session closed...")
		return false
	} else if strings.HasPrefix(str, "@shutdown") {
		fmt.Println("server shutdown...")
		exitChan <- 0
		return false
	}
	fmt.Println(str)
	return true
}