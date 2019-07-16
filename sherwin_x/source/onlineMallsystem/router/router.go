package router

import (
	"github.com/gin-gonic/gin"
	"onlineMallsystem/controller"
	"onlineMallsystem/utils"
)

func Router()  {
	r:=gin.New()

	r.POST("/sign",controller.SignHandler)

	buyer:=r.Group("/buyer")
	buyer.POST("/login", utils.GetToken("buyer").LoginHandler)
	buyer.Use(utils.GetToken("buyer").MiddlewareFunc())
	{
		buyer.POST("/hello", utils.HelloHandler)
	}

	seller:=r.Group("/seller")
	seller.POST("/login", utils.GetToken("seller").LoginHandler)
	seller.Use(utils.GetToken("seller").MiddlewareFunc())
	{
		seller.POST("/hello", utils.HelloHandler)
	}

	_ = r.Run(":8080")
}