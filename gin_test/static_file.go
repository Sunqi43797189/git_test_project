package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main()  {
	r := gin.Default()
	r.StaticFS("/myFile", http.Dir("/Users/sunqi/Desktop"))
	r.Run(":8080")
}
