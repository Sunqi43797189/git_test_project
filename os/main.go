package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.Mkdir("test_dir/test", 0777)
	if err != nil {
		fmt.Println(err)
	}
	os.MkdirAll("test_dir/test_1/inside/", 0777)
	//err := os.Remove("test_dir")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//err := os.RemoveAll("test_dir")
	//if err != nil {
	//	fmt.Println(err)
	//}
	var file *os.File
	file, err = os.Create("test_dir/test_1/inside/1.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(file)

	file, err = os.Open("test_dir/test_1/inside/1.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(file)

	file, err = os.OpenFile("test_dir/test_1/inside/1.txt", 2, 0777)

	//num, err := file.WriteString("helloworld")
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(num)

	num, err := file.Write([]byte{65})
	fmt.Println(num)


}
