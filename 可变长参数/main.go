// package main

// import "bytes"

// import "fmt"

// func main()  {
// 	res := joinString("show", "me", "your", "code")
// 	fmt.Println(res)
// }

// func joinString(strList ...string) string {
// 	var b bytes.Buffer
// 	fmt.Println(b.String())

// 	for _, str := range(strList){
// 		b.WriteString(str + " ")
// 	}
// 	return b.String()
// }


package main

func main()  {
}

func switchType(options ...interface{}) string {
	strType = ""
	for _, option := range(options){
		switch option.(type) {
		case bool:
			strType "bool"
		case string:
			strType "string"
		case int:
			strType "int"
		case float64:
			strType "float"
		default:
			strType "default"
		}
	}
	return strType
}