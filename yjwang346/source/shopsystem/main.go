package main

import (
	"github.com/gin-gonic/gin"
	"shopsystem/controller"
	"shopsystem/middleware"
)

func main(){
	router :=gin.New()
	authMiddleware := middleware.MiddleWareInit()
	user := router.Group("/user")

	router.POST("/register",controller.Signup)	//进行注册

	user.POST("login",authMiddleware.LoginHandler)
	user.Use(authMiddleware.MiddlewareFunc())
	{
		user.POST("/hello",middleware.HelloHandler)//废物功能


	}


	router.Run(":8080")
}