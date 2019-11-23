package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main(){
	router := gin.Default()
	router.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		fmt.Println(file.Filename)
		c.SaveUploadedFile(file, file.Filename)
		c.String(200, "upload success")
	})
	http.ListenAndServe(":8080", router)
}
