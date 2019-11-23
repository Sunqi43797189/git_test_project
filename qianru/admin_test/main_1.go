package admin_test

import "fmt"

type user struct {
	Name string
}

type Admin struct {
	user
}

func init(){
	fmt.Println("init")
}

var Admin1 Admin = Admin{user{Name:"sunqi"}}

func GetAdmin() string {
	return "222"
}
//func main(){
//	fmt.Println(111)
//}