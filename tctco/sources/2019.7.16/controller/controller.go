package controller

import (
	"github.com/gin-gonic/gin"
	"regexp"
)



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
	if len(password) < 6 {
		return false
	}
	return true
}
