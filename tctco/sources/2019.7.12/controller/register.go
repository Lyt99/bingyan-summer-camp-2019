package controller

import (
	"../model"
	"github.com/gin-gonic/gin"
	"net/http"
)


func Register(c *gin.Context){
	username := c.PostForm("username")
	password := c.PostForm("password")
	phonenumber := c.PostForm("phonenumber")
	email := c.PostForm("email")

	if !format_check_email(email) {
		c.JSON(400, gin.H{
			"info": "incorrect email",
		})
		return;
	}

	if !format_check_phonenumber(phonenumber) {
		c.JSON(400, gin.H{
			"info": "incorrect phonenumber",
		})
		return;
	}

	if !formatCheckPassword(password) {
		c.JSON(400, gin.H{
			"info": "password is too short!",
		})
		return;
	}

	if model.DB_isExist(username){
		c.JSON(http.StatusOK, gin.H{"info": "already registered"})
	} else {
		model.DB_register(username, password, phonenumber, email,0)
		c.JSON(http.StatusOK, gin.H{"info": "register succeed!"})
	}
}


func AdminRegister(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	phonenumber := c.PostForm("phonenumber")
	email := c.PostForm("email")

	if model.DB_isExist(username){
		c.JSON(http.StatusOK, gin.H{"info": "already registered"})
		c.Redirect(http.StatusFound, "/login")
	} else {
		model.DB_register(username, password, phonenumber, email,1)
		c.JSON(http.StatusOK, gin.H{"info": "admin register succeed!"})
	}
} // can be the same with Register
