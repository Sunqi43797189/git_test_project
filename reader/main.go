package main

import (
	"bufio"
	"os"
)

func main(){

}

func BufioRead(name string) {
	if fileObj, err := os.Open(name); err == nil {
		defer fileObj.Close()
		reader := bufio.NewReader(fileObj)
		if result, err := reader.ReadString(byte('@'));err == nil {

		}
	}
}