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
//获得自己的信息
func Getme(c *gin.Context){
	//err,user :=database.Checksignup(c)
	//那是怎么知道用户名的？
	//应该不是这样用的
	id,err := c.Get("id")
	username := id.(string)
	if !err {
		c.JSON(http.StatusBadRequest, gin.H{
			"success":false,
			"error": "获得用户id出现错误",
			"data":"",//失败的时候留空
		})
		return
	}
	error,user := database.Checksignup(username)
	if error==0{
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
//获得其他用户的信息
func Getuser(c *gin.Context){
	username :=c.Param("id")
	err,user:=database.Checksignup(username)
	if err==0{
		c.JSON(http.StatusBadRequest, gin.H{
			"success":false,
			"error": "该用户不存在Getuser",
			"data":"",//失败的时候留空
		})
		return
	}
	//这里语法不知道该怎么表示
	user.Visittime++
	c.JSON(http.StatusOK,gin.H{
		"success":true,
		"error":"",
		"data":  gin.H{
			"nickname":user.Nickname,
			"email":user.Email,
			"total_view_count":user.Visittime,
			"total_collect_count":user.Collcectcount,
		},
	})
	//还要将次数写入数据库中才行
	database.Addvisttime(c,username,user.Visittime)
}

//更改自己的信息
func Changeme(c *gin.Context){
	id,err := c.Get("id")
	username := id.(string)
	if !err {
		c.JSON(http.StatusBadRequest, gin.H{
			"success":false,
			"error": "获得用户id出现错误",
			"data":"",//失败的时候留空
		})
		return
	}
	result:=database.Changeinfo(c,username)
	if result==0{
		c.JSON(http.StatusBadRequest, gin.H{
			"success":true,
			"error": "",
			"data":"信息更改完成",//失败的时候留空
		})
	}
}
func Getmycommodity(c *gin.Context){
	id,err := c.Get("id")
	username := id.(string)
	if !err {
		c.JSON(http.StatusBadRequest, gin.H{
			"success":false,
			"error": "获得用户id出现错误",
			"data":"",//失败的时候留空
		})
		return
	}
	//username := c.Query("username")
	database.GetmeCommodity(c,username)
}