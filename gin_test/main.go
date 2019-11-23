package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
)

func main(){
	//写日志文件
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message":"pong",
		})
	})

	//r.GET("/users/:name", func(c *gin.Context) {
	//	name := c.Param("name")
	//	fmt.Println(name)
	//	c.String(http.StatusOK, "hello " + name)
	//})

	// get请求参数
	r.GET("/users/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		fmt.Println(name, action)
		c.String(http.StatusOK, name + action)
	})

	r.GET("/users/:name", func(c *gin.Context) {
		name := c.Param("name")
		firstName := c.DefaultQuery("firstname","guest")
		age := c.DefaultQuery("age", "10")
		fmt.Println(name, firstName, age)
		c.String(http.StatusOK, name + firstName + age)
	})

	//post请求表单参数
	r.POST("/users", func(c *gin.Context) {
		//必须设置请求中Content-Type 为 application/x-www-form-urlencoded
		message := c.PostForm("message")
		nickName := c.DefaultPostForm("nickname", "no_name")
		fmt.Println(message, nickName)
		c.JSON(200, gin.H{
			"message": message,
			"name": nickName,
		})
	})

	//get + post请求参数
	r.POST("/members", func(c *gin.Context) {
		id := c.Query("id") // post 请求中参数
		page := c.DefaultQuery("page", "1")
		name := c.PostForm("name") // 表单请求中的参数
		nickname := c.DefaultPostForm("nickname", "no_name")
		c.JSON(http.StatusOK, gin.H{
			"id": id,
			"page": page,
			"name": name,
			"nickname": nickname,
		})
	})

	//文件上传
	r.POST("/upload", func(c *gin.Context) {
		file, _ := c.FormFile("file")
		log.Println(file.Filename)
		c.SaveUploadedFile(file, "/Users/sunqi/Desktop")
		c.JSON(200, gin.H{
			"message": "success" + file.Filename,
		})
	})

	//多文件上传
	r.MaxMultipartMemory = 8 << 20  // 8 MiB
	r.POST("/multi-upload", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]
		for _, file := range files{
			log.Println(file)
		}
		c.JSON(200, gin.H{
			"message": "success",
		})
	})

	//路由分组
	v1 := r.Group("/v1")
	{
		v1.GET("/login")
		v1.GET("/logout")
	}

	v2 := r.Group("/v2")
	{
		v2.POST("/register")
		v2.GET("/find")
	}

	//重定向到外部
	r.GET("/redirect_outside", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	})

	//重定向到内部
	r.GET("/redirect_inside", func(c *gin.Context) {
		c.Request.URL.Path = "/inside_test"
		r.HandleContext(c)
	})

	//重定向内部测试链接
	r.GET("/inside_test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "success",
		})
	})

	r.Run("127.0.0.1:8081")
}
