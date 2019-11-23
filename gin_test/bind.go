package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type login struct {
	User     string `form:"username" json:"user" uri:"user" xml:"user"  binding:"required"`
	Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}


func main(){
	router := gin.Default()
	router.POST("/loginJson", func(c *gin.Context) {
		var login login
		if err := c.ShouldBindJSON(&login); err != nil {
			c.JSON(400, gin.H{"error":err.Error()})
			return
		}
		if login.User != "sunqi" || login.Password != "sunqi123"{
			c.JSON(403, gin.H{
				"status": "unauthorized user",
			})
			return
		}
		c.JSON(200, gin.H{
			"status": "login success",
		})
	})

	//router.GET("/:user/:password", func(c *gin.Context) {
	//	var login login
	//	if err := c.ShouldBindUri(&login); err != nil{
	//		c.JSON(200, gin.H{
	//			"status": err.Error(),
	//		})
	//		return
	//	}
	//
	//	if login.User != "sunqi" || login.Password != "sunqi123" {
	//		c.JSON(http.StatusUnauthorized, gin.H{
	//			"status": "StatusUnauthorized",
	//		})
	//		return
	//	}
	//	c.JSON(200, gin.H{
	//		"status": "login success",
	//	})
	//})
	router.StaticFS("/files", http.Dir("/Users/sunqi/go/src"))
	//router.Static("/go_files", "/User/sunqi/go/src")
	http.ListenAndServe(":8080", router)
}

