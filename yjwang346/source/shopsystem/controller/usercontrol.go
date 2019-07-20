package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shopsystem/database"
	"shopsystem/model"
)

func Login(c *gin.Context){
	var user model.Register
	//newuser.UserName = c.PostForm("username")
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success":false,
			"error": "传递JSON参数出现错误",
			"data":"",//失败的时候留空
		})
		return
	}
}

func Getme(c *gin.Context){
	//err,user :=database.Checksignup(c)
	//那是怎么知道用户名的？
	//应该不是这样用的
}

func Getuser(c *gin.Context){
	username :=c.Param("id")
	err,user:=database.Checksignup(username)
	if err==0{
		c.JSON(http.StatusBadRequest, gin.H{
			"success":false,
			"error": "该用户不存在",
			"data":"",//失败的时候留空
		})
		return
	}
	//这里语法不知道该怎么表示
	c.JSON(http.StatusOK,gin.H{
		"success":true,
		"error":"",
		"data":  gin.H{
			"nickname":user.Nickname,
			"email":user.Email,
			"total_view_count":user.Visittime,
			"total_collect_count":user,
		},
	})
}