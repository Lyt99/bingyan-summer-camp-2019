package controller

import (
	"../model"
	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Register(c *gin.Context){
	username := c.PostForm("username")
	password := c.PostForm("password")
	phonenumber := c.PostForm("phonenumber")
	email := c.PostForm("email")

	if model.DB_isExist(username){
		c.JSON(http.StatusOK, gin.H{"info": "already registered"})
		c.Redirect(http.StatusFound, "/login")
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


func Login(c *gin.Context) (interface{}, error) { //authCallback?
	username := c.PostForm("username")
	password := c.PostForm("password")

	result := model.DB_search_user(username)
	switch user := result.(type){
	case model.User:
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if err == nil {
			c.JSON(http.StatusOK, gin.H{"info": "login succeed!",
				"ID":        user.ID,
				"username":  user.Username,
				"password":  user.Password,
				"authority": user.Authority,
			})
			return user, nil
		} else {
			c.JSON(http.StatusOK, gin.H{"info": "wrong password!"})
		}
	default:
		c.JSON(http.StatusOK, gin.H{"info": "no such user!"})
	}
	return nil, nil
}


func AdminAuthCheck(data interface{}, c *gin.Context) bool { // authPrivCallback
	if user, ok := data.(*model.User); ok && user.Authority > 0 {
		c.JSON(200, gin.H{
			"info": "you are an admin!",
		})
		return true
	} else {
		c.JSON(403, gin.H{
			"info": "you have no right!",
		})
		return false
	}
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

//func AdminManage(c *gin.Context) {
//	claims := jwt.ExtractClaims(c)
//	c.JSON(200, gin.H{
//		""
//	})
//}