package controller

import (
	"../model"
	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

// auth fail
func UnAuthFunc(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code": code,
		"message": message,
	})
}

// fill in the payload of jwt
func FillPayload(data interface{}) jwt.MapClaims {
	if user, ok := data.(*model.User); ok {
		return jwt.MapClaims{
			"username": user.Username,
			"authority": user.Authority,
		}
	}
	return jwt.MapClaims{}
}

//Create this middle ware
func MiddleWareInit () *jwt.GinJWTMiddleware  {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm: "You have to login!",
		Key:[]byte("secret key salt"),
		Timeout:time.Hour,
		MaxRefresh:time.Hour,
		IdentityKey: "username",
		PayloadFunc:FillPayload,
		Authenticator: Login, // authentication
		Unauthorized:UnAuthFunc,
		TokenLookup:"header: Authorization, query: token, cookie: jwt",
		TokenHeadName:"Bearer",
		TimeFunc:time.Now,
	})
	if err != nil{
		log.Fatal("JWT Error:"+err.Error())
	}
	return authMiddleware
}

