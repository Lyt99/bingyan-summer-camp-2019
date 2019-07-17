package controller

import (
	"demo/model"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	register(0, c)
}

func AdminRegister(c *gin.Context) { // can be the same with Register
	register(1, c)
}

func register(authority int, c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	phonenumber := c.PostForm("phonenumber")
	email := c.PostForm("email")

	if !format_check_email(email) {
		c.JSON(400, gin.H{
			"info": "incorrect email",
		})
		return
	}

	if !format_check_phonenumber(phonenumber) {
		c.JSON(400, gin.H{
			"info": "incorrect phonenumber",
		})
		return
	}

	if !formatCheckPassword(password) {
		c.JSON(400, gin.H{
			"info": "password is too short!",
		})
		return
	}

	if exist, err := model.DB_isExist(username); exist {
		c.JSON(400, gin.H{"info": "already registered"})
		return
	} else if err != nil {
		c.JSON(500, gin.H{"info": "somehow failed"})
	}

	if model.DB_register(username, password, phonenumber, email, authority) {
		c.JSON(200, gin.H{"info": "register succeed!"})
	} else {
		c.JSON(500, gin.H{"info": "somehow failed"})
	}

}
