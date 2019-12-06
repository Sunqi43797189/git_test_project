package main

import "fmt"

type Service interface {
	Start()
	Log(string)
}

type Logger struct {
}

func (g *Logger) Log(i  string) {
	fmt.Println(i)
}

type GameService struct {
	Logger
}

func (g *GameService)Start()  {
	fmt.Println("game start")
}

func main()  {
	var s Service = new(GameService)
	s.Start()
	s.Log("hello")	
}