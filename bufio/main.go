package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main(){
	data := []byte("中华人民共和国")
	rd := bytes.NewReader(data)
	fmt.Println(rd)

	r := bufio.NewReader(rd)
	//fmt.Println(r)

	var buf [128]byte
	n, err := r.Read(buf[:])
	fmt.Println(string(buf[:n]), n,  err)
}
