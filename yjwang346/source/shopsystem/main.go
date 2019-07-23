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
	me := router.Group("/me")
	commodity := router.Group("/commodity")

	router.POST("/register",controller.Signup)	//进行注册
	user.POST("/login",authMiddleware.LoginHandler)

	//试验用
	//router.POST("/login",database.Checklogin2)

	me.Use(authMiddleware.MiddlewareFunc())
	{
		me.GET("",controller.Getme)	//获得自己的信息
		me.POST("",controller.Changeme)	//更改自己的信息
	}


	user.Use(authMiddleware.MiddlewareFunc())
	{
		user.POST("/hello",middleware.HelloHandler)//废物功能
		user.GET("/:id",controller.Getuser)	//获得其他用户的相关信息
	}

	commodity.Use(authMiddleware.MiddlewareFunc())
	{
		commodity.POST("",controller.Postcommodity)
	}

	router.Run(":8080")
}