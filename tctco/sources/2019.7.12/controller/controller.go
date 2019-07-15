package controller

import (
	"../model"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"regexp"
)


func Login(c *gin.Context) (interface{}, error) { //authCallback?
	username := c.PostForm("username")
	password := c.PostForm("password")

	result := model.DB_search_user(username)
	if user, ok := result.(*model.User); ok{
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"info": "login succeed!",
				"ID":        user.ID,
				"username":  user.Username,
				"password":  user.Password,
				"authority": user.Authority,
			})
			return user, nil
		} else {
			c.JSON(http.StatusOK, gin.H{"info": "wrong password!"})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{"info": "no such user!"})
	}
	return nil, nil
}


func MainPage(c *gin.Context) {
	c.JSON(200, gin.H{
		"info": "hello world!",
	})
}

// check format
func format_check_email(email string) bool {
	matched, _ := regexp.MatchString(`^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z]+)+$`, email)
	return matched
}


func format_check_phonenumber(phonenumber string) bool {
	matched, _ := regexp.MatchString(`^((13[0-9])|(15[^4])|(18[0,2,3,5-9])|(17[0-8])|147)\d{8}$`, phonenumber)
	return matched
}


func formatCheckPassword(password string) bool {
	if len(password) < 6{
		return false
	}
	return true
}