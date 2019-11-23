package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Login struct {
	username string `json:"username" binding: "required"`
	password string `json:"password" binding: "required"`
}

func v1Login(c *gin.Context){
	var login Login

	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if login.username != "sunqi" || login.password != "sunqi" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": "unauthorized",
		})
		return
	}
	c.JSON(200, gin.H{
		"status": "login success",
	})
}


func main(){
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.POST("/login_json", v1Login)
	}

	r.Run()
}
