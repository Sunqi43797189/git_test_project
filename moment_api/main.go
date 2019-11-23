package main

import "go_test_project/moment_api/app/routers"

func main(){
	routers := routers.InitRouter()
	routers.Run(":8081")
}