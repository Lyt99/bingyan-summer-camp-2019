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
	r.POST("/register", controller.Register)
	r.POST("/register/admin", controller.AdminRegister)
	//r.POST("/login", controller.Login)


	//authorized := r.Group("/admin", controller.AuthHandler)
	//authorized.GET("/manage", controller.ManageHandler)

	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm: "You have to login!",
		Key:[]byte("secret key salt"),
		Timeout:time.Hour,
		MaxRefresh:time.Hour,
		IdentityKey: "username",
		PayloadFunc:controller.FillPayload,
		Authenticator: controller.Login,
		Authorizator:controller.AdminAuthCheck,
		Unauthorized:controller.UnAuthFunc,
		TokenLookup:"header: Authorization, query: token, cookie: jwt",
		TokenHeadName:"Bearer",
		TimeFunc:time.Now,
	})
	if err != nil{
		log.Fatal("JWT Error:"+err.Error())
	}


	r.POST("/login", authMiddleware.LoginHandler)
	auth := r.Group("/admin")
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		auth.GET("/manege", )
	}

	r.Run(":8080")
}
