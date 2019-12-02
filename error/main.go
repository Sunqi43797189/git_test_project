package main

import "errors"

import "fmt"

// import "fmt"

// func New(text string) (e *errorString) {
// 	return &errorString{text}
// }

// type errorString struct {
// 	s string
// }

// func (e *errorString)Error() string {
// 	return e.s
// }

// func main() {
// 	e1 := New("new error")
// 	fmt.Println(e1.Error())
// }
var err = errors.New("除数为零")

func main() {
	res, err := div(1,0)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(res)
}

func div(beichushu, chushu int) (int, error) {
	if chushu == 0 {
		return 0, err
	} else {
		return beichushu / chushu, nil
	}
}