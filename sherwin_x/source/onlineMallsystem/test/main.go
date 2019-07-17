package main

import (
	"github.com/gin-gonic/gin"
	"onlineMallsystem/controller"
	"onlineMallsystem/utils"
)

func main() {
	r := gin.Default()

	r.POST("/user", controller.SignIn)                   //注册
	r.POST("/user/login", utils.GetToken().LoginHandler) //登录

	user:=r.Group("/user")
	user.Use(utils.GetToken().MiddlewareFunc())
	user.GET("/:id",controller.ShowUser) 					//查看某位用户资料

	me := r.Group("/me")
	me.Use(utils.GetToken().MiddlewareFunc())
	me.GET("/hello", utils.HelloHandler)                   //测试
	me.GET("", controller.ShowMe)                          //查看个人资料
	//me.POST("", controller.UpdateMe)                       //修改个人资料
	me.GET("/commodities", controller.MyCommodities)       //查看我的发布
	//me.GET("/collections", controller.MyCollections)       //查看我的收藏
	//me.POST("/collections", controller.NewCollection)      //收藏某个商品
	//me.DELETE("/collections", controller.DeleteCollection) //删除某个收藏

	commodities := r.Group("/commodities")
	commodities.Use(utils.GetToken().MiddlewareFunc())
	//commodities.GET("", controller.GetCommodities)   //获取商品列表
	//commodities.GET("/hot", controller.GetHotSearch) //获取热搜词
	commodities.POST("", controller.NewCommodity)    //发布新商品

	commodity := r.Group("/commodity")
	commodity.Use(utils.GetToken().MiddlewareFunc())
	commodity.GET("/:id", controller.DetailCommodity)    //某个商品详情
	commodity.DELETE("/:id", controller.DeleteCommodity) //删除某个商品

	_ = r.Run(":8080")
}
