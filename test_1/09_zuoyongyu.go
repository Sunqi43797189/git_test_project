package main

import (
	"log"
	"os"
	"fmt"
)


var cwd string

func init() {
    cwd, err := os.Getwd() // NOTE: wrong!
    if err != nil {
        log.Fatalf("os.Getwd failed: %v", err)
    }
    log.Printf("Working directory = %s", cwd)
}

func zuoYongYuTest() {
  fmt.Println("finish")
  fmt.Println(cwd)
}