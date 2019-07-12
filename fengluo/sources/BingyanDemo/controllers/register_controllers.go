package controllers

import (
	"BingyanDemo/models"
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterPost(c *gin.Context)  {
	username := c.DefaultPostForm("username","")
	password := c.DefaultPostForm("password", "")
	userid := c.DefaultPostForm("userid", "")
	userphone := c.DefaultPostForm("userphone", "")
	useremail := c.DefaultPostForm("useremail","")
	fmt.Println(username, password, userid, userphone, useremail)
	id := models.IfNameRepeat(username)
	fmt.Println("id:",id)
	if id > 0 {
		c.JSON(http.StatusOK, gin.H{"message": "用户昵称已存在"})
		return
	}
	id =models.IfIdRepeat(userid)
	fmt.Println("id:",id)
	if id > 0 {
		c.JSON(http.StatusOK, gin.H{"message": "用户ID已存在"})
		return
	}
	if models.IfNil(username, userphone, userid, useremail, password) > 0 {
		c.JSON(http.StatusOK, gin.H{"message": "信息未完善"})
		return
	}
	password = fmt.Sprintf("%x",md5.Sum([]byte(password)))
	user :=models.User{userid, password, username, false, userphone, useremail}
	_,err :=models.InsertUser(user)
	if err != nil{
		c.JSON(http.StatusOK, gin.H{"message": "注册失败"})
	}else {
		c.JSON(http.StatusOK, gin.H{"message": "注册成功"})
	}
}