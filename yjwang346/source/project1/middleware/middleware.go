package middleware

import (
	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"log"
	"project1/controllers"
	"time"
)
//Create this middle ware
func UserMiddleWareInit () *jwt.GinJWTMiddleware  {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm: "You have to login!",
		Key:[]byte("secret key salt"),
		Timeout:time.Hour,
		MaxRefresh:time.Hour,
		IdentityKey: "username",
		// Callback function that should perform the authentication of the user based on userID and
		// password. Must return true on success, false on failure. Required.
		// Option return user data, if so, user data will be stored in Claim Array.
		//必要项, 这个函数用来判断 User 信息是否合法，如果合法就反馈 true，否则就是 false, 认证的逻辑就在这里
		Authenticator: controllers.LoginPost, // authentication
		//Unauthorized:UnAuthFunc,
		//整体上这次编程并没有考虑权限的因素
		TokenLookup:"header: Authorization, query: token, cookie: jwt",
		TokenHeadName:"Bearer",
		TimeFunc:time.Now,
	})
	if err != nil{
		log.Fatal("JWT Error:"+err.Error())
	}
	return authMiddleware
}
func AdminMiddleWareInit () *jwt.GinJWTMiddleware  {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm: "You have to login!",
		Key:[]byte("secret key salt"),
		Timeout:time.Hour,
		MaxRefresh:time.Hour,
		IdentityKey: "username",
		// Callback function that should perform the authentication of the user based on userID and
		// password. Must return true on success, false on failure. Required.
		// Option return user data, if so, user data will be stored in Claim Array.
		//必要项, 这个函数用来判断 User 信息是否合法，如果合法就反馈 true，否则就是 false, 认证的逻辑就在这里
		Authenticator: controllers.Checkadmin, // authentication
		//Unauthorized:UnAuthFunc,
		//整体上这次编程并没有考虑权限的因素
		TokenLookup:"header: Authorization, query: token, cookie: jwt",
		TokenHeadName:"Bearer",
		TimeFunc:time.Now,
	})
	if err != nil{
		log.Fatal("JWT Error:"+err.Error())
	}
	return authMiddleware
}