package main
import (
	"fmt"
	"go_test_project/qianru/admin_test"
	)
//type Hello interface {
//	hello()
//}
//
//type User struct {
//	name string
//	email string
//}
//
//func (u User) hello() {
//	fmt.Println(111)
//}
//
//func sayHello(h Hello){
//	h.hello()
//}
//
func main(){
	//user := User{name: "sunqi", email: "111"}
	//sayHello(user)
	fmt.Println(admin_test.GetAdmin())
}