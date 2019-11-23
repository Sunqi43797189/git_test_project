package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func uploadFile(c *gin.Context){
	file, err := c.FormFile("file")
	if err != nil {
		log.Println(err.Error())
		c.String(500, "upload fail")
	}
	log.Println(file.Filename)
	c.SaveUploadedFile(file, file.Filename)
	c.String(200, "upload success")
}
func main(){
	router := gin.Default()
	router.GET("/hello-world", func(c *gin.Context) {
		c.String(200, "hello gin")
	})

	router.POST("/upload_image", uploadFile)
	http.ListenAndServe("127.0.0.1:8080", router)
	//router.Run("127.0.0.1:8080")
}


