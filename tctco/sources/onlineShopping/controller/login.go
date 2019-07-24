package controller

import (
	"fmt"
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"onlineShopping/model"
)


func Hello(c *gin.Context) {
	c.JSON(200, gin.H{"info": "hello"})
}

func Login(c *gin.Context) (interface{}, error) {
	var json model.LoginJSON
	response := ResponseInit()

	if err := c.ShouldBindJSON(&json); err != nil {
		response["error"] = "输入信息不完整"
		c.JSON(400, response)
		return nil, jwt.ErrFailedAuthentication
	}

	exist, err := model.DBIsUserExist(json.Username)
	if err != nil {
		dbError(c)
		return nil, jwt.ErrFailedAuthentication
	}

	if !exist {
		response["error"] = "用户未注册"
		c.JSON(400, response)
		return nil, jwt.ErrFailedAuthentication
	}

	user, err := model.DBSearchUser(json.Username)
	if err != nil {
		dbError(c)
		return nil, jwt.ErrFailedAuthentication
	}

	if !checkPassword(json.Password, user.Password) {
		response["error"] = "密码不正确"
		c.JSON(400, response)
		return nil, jwt.ErrFailedAuthentication
	}
	return user, nil
}

func checkPassword(userPassword, dbPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(userPassword))
	fmt.Println(dbPassword)
	if err != nil {
		return false
	}
	return true
}
