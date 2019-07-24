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
	commodities := router.Group("/commodities")

	router.POST("/register",controller.Signup)	//进行注册
	user.POST("/login",authMiddleware.LoginHandler)

	//试验用
	//router.POST("/login",database.Checklogin2)

	me.Use(authMiddleware.MiddlewareFunc())
	{
		me.GET("",controller.Getme)	//获得自己的信息
		me.POST("",controller.Changeme)	//更改自己的信息
		me.GET("/commodities",controller.Getmycommodity)
	}
	//router.GET("/commodities",controller.Getmycommodity)

	user.Use(authMiddleware.MiddlewareFunc())
	{
		user.POST("/hello",middleware.HelloHandler)//废物功能
		user.GET("/:id",controller.Getuser)	//获得其他用户的相关信息
	}

	commodity.Use(authMiddleware.MiddlewareFunc())
	{
		//获得热门查询关键词，可能有问题
		commodity.POST("hot",controller.Get_hot_keyword)

		//获得单个商品的信息
		commodity.GET("/:id",controller.Getone_commodityinfo)
	}

	commodities.Use(authMiddleware.MiddlewareFunc())
	{
		//用户发布商品
		commodities.POST("",controller.Postcommodity)
		//获得商品列表
		commodities.GET("",controller.Getcommodities)
	}

	//commodity.GET("/:id",controller.Getone_commodityinfo)
	router.POST("/searchword",controller.Get_hot_keyword)
	//？？router.POST("/pics",database.Picture)
	//关于这个上传图片这个功能不是很明白

	router.Run(":8080")
}