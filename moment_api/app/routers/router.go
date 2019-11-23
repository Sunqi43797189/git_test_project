package routers

import (
	"github.com/gin-gonic/gin"
	v1 "go_test_project/moment_api/app/routers/v1"
)

func InitRouter() *gin.Engine{
	router := gin.Default()
	api_v1 := router.Group("/v1")
	{
		api_v1.GET("/moments", v1.GetAllMoments)
		api_v1.GET("/moments/:id", v1.GetMomentsById)
		api_v1.POST("/moments/:moment_id", v1.UpdateMomentsById)
		api_v1.DELETE("/moments/:moment_id", v1.DeleteMomentById)
	}

	return router
}
