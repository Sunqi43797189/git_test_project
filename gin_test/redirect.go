package main

import "github.com/gin-gonic/gin"

func main(){
	r := gin.Default()

	r.GET("/redirect_test", func(c *gin.Context) {
		c.Redirect(301, "http://www.baidu.com")
	})

	r.Run(":8080")
}
