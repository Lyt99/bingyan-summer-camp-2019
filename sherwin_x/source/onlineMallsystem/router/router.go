package router

import (
	"github.com/gin-gonic/gin"
	"onlineMallsystem/controller"
	"onlineMallsystem/utils"
)

func Router() {
	r := gin.New()

	r.POST("/user", controller.SignHandler)
	r.POST("/user/login", utils.GetToken().LoginHandler)
	//r.GET(" /user/:id")

	me := r.Group("/me")
	me.Use(utils.GetToken().MiddlewareFunc())
	{
		me.GET("/hello", utils.HelloHandler)
		me.GET("", controller.ShowHandler)
	}

	commodities := r.Group("/commodities")
	commodities.Use(utils.GetToken().MiddlewareFunc())
	{
		commodities.GET("/hello", utils.HelloHandler)
		//commodities.GET("")
		//commodities.GET("/hot")
		commodities.POST("",controller.NewCommodity)
		//commodities.GET("/:id")
		//commodities.DELETE("/:id")
	}

	_ = r.Run(":8080")
}
