package router

import (
	"github.com/gin-gonic/gin"
	"onlineMallsystem/controller"
	"onlineMallsystem/utils"
)

func Router() {
	r := gin.New()

	r.POST("/user", controller.SignIn)
	r.POST("/user/login", utils.GetToken().LoginHandler)
	//r.GET(" /user/:id")

	me := r.Group("/me")
	me.Use(utils.GetToken().MiddlewareFunc())
	{
		//me.GET("/hello", utils.HelloHandler)
		me.GET("", controller.ShowMe)
		//me.POST("")
		me.GET("/commodities",controller.MyCommodities)
		//me.GET("/collections")
		//me.POST("/collections")
		//me.DELETE("/collections")
	}

	commodities := r.Group("/commodities")
	commodities.Use(utils.GetToken().MiddlewareFunc())
	{
		//commodities.GET("")
		//commodities.GET("/hot")
		commodities.POST("",controller.NewCommodity)
	}

	commodity:=r.Group("/commodity")
	commodity.Use(utils.GetToken().MiddlewareFunc())
	{
		commodity.GET("/:id",controller.DetailCommodity)
		//commodity.DELETE("/:id")
	}

	_ = r.Run(":8080")
}
