package controllers

import (
	"BingyanDemo/models"
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginPost(c *gin.Context) {
	username := c.DefaultPostForm("username","")
	password := c.DefaultPostForm("password", "")
	password = fmt.Sprintf("%x",md5.Sum([]byte(password)))
	fmt.Println("username:", username, ",password:", password)

	id := models.LoginCheck(username,password)
	fmt.Println("id:", id)
	if id > 0 {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "登录成功"})

	} else {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "登录失败"})
	}
}
