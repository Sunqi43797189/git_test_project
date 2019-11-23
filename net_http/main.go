package main

import (
	"fmt"
	"net"
)

func main() {
    //http.ListenAndServe("127.0.0.1:8080", nil)
    ip, err := net.ResolveIPAddr("ip", "www.google.com")
    if err != nil {
    	fmt.Println(err)
	}
    fmt.Println(ip)
}