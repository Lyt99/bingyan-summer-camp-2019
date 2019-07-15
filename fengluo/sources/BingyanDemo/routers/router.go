package routers

import (
	"BingyanDemo/controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	router := gin.Default()
	router.POST("/register", controllers.RegisterPost)
	router.POST("/login", controllers.LoginPost)

	return router

}
