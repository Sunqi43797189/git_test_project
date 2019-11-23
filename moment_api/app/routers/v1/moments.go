package v1

import (
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"go_test_project/moment_api/models"
	"time"
)

func GetAllMoments(c *gin.Context){
	//返回全部评论 /api/v1/moments
	data := make(map[string]interface{})
	data["count"] = models.GetMomentsCount()
	data["lists"] = models.GetAllMoments()
    c.JSON(200, gin.H{
    	"message": "success",
    	"data": data,
	})
}

func GetMomentsById(c *gin.Context){
    data := make(map[string]interface{})
    id := c.Param("id")
    data["lists"] = models.GetMomentById(id)
    c.JSON(200, gin.H{
    	"message": "success",
    	"data": data,
	})
}

func UpdateMomentsById(c *gin.Context){
	id, _ := com.StrTo(c.Param("moment_id")).Int()
	memberId := c.Query("member_id")
	likeCount, _ := com.StrTo(c.Query("like_count")).Int()
	content := c.Query("content")
	status, _ := com.StrTo(c.Query("status")).Int()
    if id == 0 {
		c.JSON(200, gin.H{
			"message": "fail",
			"error": "id 不能为空",
		})
	} else {
		valid := validation.Validation{}
		valid.Required(memberId, "member_id").Message("member_id 不能为空")
		valid.Required(likeCount, "like_count").Message("like_count 不能为空")
		valid.Range(status, 0, 1, "status").Message("status 只能为0或1")
		valid.Required(content, "content").Message("content 不能为空")
		if !valid.HasErrors(){
			maps := make(map[string]interface{})
			maps["member_d"] = memberId
			maps["like_count"] = likeCount
			maps["content"] = content
			maps["status"] = status
			maps["updated_at"] = time.Now()
			if models.MomentExistsById(id){
				res := models.UpdateMomentsById(id, maps)
				if res{
					c.JSON(200, gin.H{
						"message": "success",
					})
				} else {
					c.JSON(500, gin.H{
						"message":"fail",
					})
				}
			} else {
				c.JSON(200, gin.H{
					"message": "fail",
					"error": "id不存在",
				})
			}
		} else {
			c.JSON(400, gin.H{
				"message": "fail",
				"error": valid.Errors,
			})
		}
	}
}

func DeleteMomentById(c *gin.Context){
	id, _ := com.StrTo(c.Param("moment_id")).Int()
	var code int = 0
	var res = make(map[string]interface{})
	if models.MomentExistsById(id) {
		models.DeleteMomentById(id)
		code = 200
		res["message"] = "success"
	} else {
		code = 400
		res["message"] = "fail"
		res["error"] = "id 不存在"
	}
	c.JSON(code, gin.H{
		"data": res,
	})
}
