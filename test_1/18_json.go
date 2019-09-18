package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title string
	Year int `json:"released"`
	Color bool  `json:"color,omitempty"`
	Actors []string
}

func jsonTest()  {
	var movies = []Movie{
		{Title: "世界", Year: 1942, Color: false,
			Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "你好", Year: 1967, Color: true,
			Actors: []string{"Paul Newman"}},
		{Title: "哪吒", Year: 1968, Color: true,
			Actors: []string{"Steve McQueen", "Jacqueline Bisset"}}}
	fmt.Println(movies)

	jsonMarshaling(movies)
}

func jsonMarshaling(movies []Movie)  {
	data, err := json.Marshal(movies)
	if err != nil {
		fmt.Println("err")
	}
	fmt.Printf("%s\n", data)

	data, err = json.MarshalIndent(movies, "", "\t")
	if err != nil {
		fmt.Println("err")
	}
	fmt.Printf("%T\n", data)
	fmt.Printf("%s\n", data)

	jsonUnMarshaling(data)
}

func jsonUnMarshaling(data []uint8) {
	var titles []struct{Title string}
	if err := json.Unmarshal(data, &titles); err != nil {
		fmt.Println("err")
	}
	fmt.Println(titles)
}