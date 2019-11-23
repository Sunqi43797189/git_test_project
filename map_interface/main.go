package main

import (
	"fmt"
)

type Dictionary struct {
	data map[interface{}]interface{}
}

func (d *Dictionary) GetKey(key interface{}) interface{} {
	return d.data[key]
}

func (d *Dictionary)SetKey(key interface{}, value interface{}) {
	d.data[key] = value
}

func (d *Dictionary)Visit(){
	for key, value := range d.data{
		fmt.Println(key, value)
	}
}

func (d *Dictionary)Clear() {
	d.data = make(map[interface{}]interface{})
}

func NewDictionary() *Dictionary {
	d := &Dictionary{}
	d.Clear()
	return d
}

func main() {
	dict := NewDictionary()
	fmt.Println(dict)
	dict.SetKey("My Factory", 60)
	dict.SetKey("Terra Craft", 36)
	dict.SetKey("Don't Hungry", 24)
	fmt.Println(dict)

	value := dict.GetKey("My Factory")
	fmt.Println(value)
	dict.Visit()

}
