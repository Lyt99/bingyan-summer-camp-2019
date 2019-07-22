package middleware

import (
	"awesomeProject/controller"
	jwt "github.com/appleboy/gin-jwt"
	"log"
	"time"
)

func GetAdminToken() *jwt.GinJWTMiddleware {
	adminTaken, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:         "test",
		Key:           []byte("sherwin"),
		Timeout:       time.Hour,
		MaxRefresh:    time.Hour,
		Authenticator: controller.AdminCallback,
		Unauthorized:  UnauthorizedFun,
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
		return nil
	} else {
		return adminTaken
	}
}
