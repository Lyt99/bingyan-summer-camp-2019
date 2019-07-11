package middleware

import (
	"awesomeProject/controller"
	"awesomeProject/model"
	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)



func GetUserToken() *jwt.GinJWTMiddleware {
	userToken,err:=jwt.New(&jwt.GinJWTMiddleware{

		Realm:"test",
		Key:[]byte("user"),
		Timeout:time.Hour,
		MaxRefresh:time.Hour,
		//put user id into token
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*model.LoginForm); ok {
				return jwt.MapClaims{
					model.IdentityKey: v.Id,
				}
			}
			return jwt.MapClaims{}
		},
		Authenticator:controller.UserCallback,
		//Authorizator: adminPrivCallback,
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		IdentityKey: "id",
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc: time.Now,
	})
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
		return nil
	}else {return userToken}
}
