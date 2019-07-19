package middleware

import (
	"awesomeProject/model"
	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func MiddleWareInit () *jwt.GinJWTMiddleware  {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm: "You have to login!",
		Key:[]byte("secret key salt"),
		Timeout:time.Hour,
		MaxRefresh:time.Hour,
		IdentityKey: "id",
		PayloadFunc:PayloadFunc,

		//用来判断用户是否是正确用户
		Authenticator: Login, // authentication
		//如果认证不成功的处理
		Unauthorized:UnAuthFunc,
		TokenLookup:"header: Authorization, query: token, cookie: jwt",
		TokenHeadName:"Bearer",
		TimeFunc:time.Now,
	})
	if err != nil{
		log.Fatal("JWT Error:"+err.Error())
		return nil
	}
	return authMiddleware
}

//put user id into token
func PayloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(*model.LoginForm); ok {
		return jwt.MapClaims{
			"id": v.Id,					//这里“id”就是代表着 IdentityKey
		}
	}
	return jwt.MapClaims{}
}

//return login failed message
func UnauthorizedFun(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}