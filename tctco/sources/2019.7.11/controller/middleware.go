package controller

import (
	"../model"
	"fmt"
	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

func AdminAuthCheck(data interface{}, c *gin.Context) bool { // authPrivCallback
	fmt.Println(data)
	if temp_user, ok := data.(model.User); ok {
		if temp_user.Authority > 0 {
			c.JSON(200, gin.H{
				"info": "you are an admin",
			})
			return true
		} else {
			c.JSON(403, gin.H{
				"info": "you are a normal user, you have no right!",
			})
		}
	} else {
		c.JSON(403, gin.H{
			"info": "you are nothing! go away!",
		})
	}
	return false
}


func UnAuthFunc(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code": code,
		"message": message,
	})
}


func FillPayload(data interface{}) jwt.MapClaims {
	if user, ok := data.(*model.User); ok {
		return jwt.MapClaims{
			"username": user.Username,
			"authority": user.Authority,
		}
	}
	return jwt.MapClaims{}
}


func IdentityTeller(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	var temp_user model.User

	username, username_ok := claims["username"].(string)
	authority, authority_ok := claims["authority"].(float64)
	if username_ok && authority_ok {
		temp_user.Username = username
		temp_user.Authority = int(authority)
		fmt.Println(temp_user.Username, temp_user.Authority)
		return temp_user
	}
	return nil
}