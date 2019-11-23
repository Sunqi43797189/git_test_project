package main

import (
	"fmt"
)

type DataWriter interface {
	WriteData(data interface{}) error
	ReadData()
}

type Service interface {
	Start()
	Log(info string)
}

type file struct {

}

type Logger struct {

}

type GameService struct {
	Logger
}

func (l *Logger)Log(info string) {
	fmt.Println(info)
}

func (g *GameService)Start(){
	fmt.Println("start game")
}

func (d *file)WriteData(data interface{}) error {
	fmt.Println("writing data")
	return nil
}

func (d *file)ReadData(){
	fmt.Println("reading data")
}

//func (d *file)WriteDataX(data interface{}) error {
//	fmt.Println("wirting data v2")
//	return nil
//}

func main(){
	//f := new(file)
	//
	//var writer DataWriter
	//writer = f //必须实现接口所有方法且方法格式与接口中的方法完全一致，才可以赋值给接口对象
	//
	//writer.WriteData("data")
	game := new(GameService) //嵌套结构体实现接口
	var service Service
	service = game
	service.Start()
	service.Log("info")
}