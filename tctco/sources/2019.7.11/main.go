package main

import (
	"./controller"
	"github.com/gin-gonic/gin"
	"github.com/appleboy/gin-jwt"
	"log"
	"time"
)


func main() {
	r := gin.Default()
	r.GET("/", controller.MainPage)
	r.POST("/register", controller.Register)
	r.POST("/register/admin", controller.AdminRegister)


	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm: "You have to login!",
		Key:[]byte("secret key salt"),
		Timeout:time.Hour,
		MaxRefresh:time.Hour,
		IdentityKey: "username",
		IdentityHandler: controller.IdentityTeller, // return username and authority
		PayloadFunc:controller.FillPayload,
		Authenticator: controller.Login, // authentication
		Authorizator:controller.AdminAuthCheck, //further check if the current user has authority
		Unauthorized:controller.UnAuthFunc,
		TokenLookup:"header: Authorization, query: token, cookie: jwt",
		TokenHeadName:"Bearer",
		TimeFunc:time.Now,
	})
	if err != nil{
		log.Fatal("JWT Error:"+err.Error())
	}


	r.POST("/login", authMiddleware.LoginHandler)
	r.GET("/:username", controller.InfoPage)
	r.POST("/:username", controller.UpdateHandler)
	auth := r.Group("/admin")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/manage/:pageno", controller.AdminManage)
	}




	r.Run(":8080")
}