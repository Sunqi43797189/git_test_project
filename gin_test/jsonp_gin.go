package main

import "github.com/gin-gonic/gin"

func main(){
	r := gin.Default()
	r.GET("/jsonp", func(c *gin.Context) {
		data := map[string]interface{}{
			"foo": "bar",
		}
		c.JSONP(200, data)
	})

	r.Run(":8080")
}
