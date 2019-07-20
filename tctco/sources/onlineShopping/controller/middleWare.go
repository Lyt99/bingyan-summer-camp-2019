package controller

import (
	"fmt"
	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"log"
	"onlineShopping/model"
	"time"
)

func MiddleWareInit() *jwt.GinJWTMiddleware {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:         "身份验证",
		Key:           []byte("secret key salt"),
		Timeout:       time.Hour,
		MaxRefresh:    time.Hour,
		IdentityKey:   "username",
		PayloadFunc:   fillPayload,
		Authenticator: Login, // authentication
		LoginResponse: loginSuccess,
		Unauthorized:  unAuthFunc,
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	return authMiddleware
}

func fillPayload(data interface{}) jwt.MapClaims {
	if user, ok := data.(model.User); ok {
		payload := make(map[string]interface{})
		payload["username"] = user.Username
		payload["id"] = user.ID
		return payload
	}
	fmt.Println("didn't fill in the payload!")
	return jwt.MapClaims{}
}

func loginSuccess(c *gin.Context, code int, token string, expire time.Time) {
	response := make(Response)
	response["success"] = true
	response["error"] = ""
	response["data"] = token
	c.JSON(200, response)
}

func unAuthFunc(c *gin.Context, code int, message string) {
	response := ResponseInit()
	fmt.Println("fail!")
	response["error"] = "需要认证"
	c.JSON(401, response)
}
