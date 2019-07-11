package middleware

import (
	"awesomeProject/controller"
	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func GetAdminToken() *jwt.GinJWTMiddleware{
	adminTaken, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:"test",
		Key:[]byte("sherwin"),
		Timeout:time.Hour,
		MaxRefresh:time.Hour,
		Authenticator:controller.AdminCallback,
		//Authorizator: adminPrivCallback,
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc: time.Now,
	})
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
		return nil
	}else {return adminTaken}
}