package view

import (
	"demo/controller"
	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
)

func RouterInit(authMiddleware *jwt.GinJWTMiddleware) *gin.Engine {
	r := gin.Default()

	r.GET("/", controller.MainPage)
	r.POST("/register", controller.Register)
	r.POST("/login", authMiddleware.LoginHandler)

	admin := r.Group("/admin")
	r.POST("/register/admin", controller.AdminRegister)
	admin.Use(authMiddleware.MiddlewareFunc())
	{
		admin.GET("/manage/:pageno", controller.AdminManage)
	}

	r.GET("/user/center/:username", controller.InfoPage)
	user := r.Group("/user")
	user.Use(authMiddleware.MiddlewareFunc())
	{
		user.POST("/center/:username", controller.UpdateHandler)
		user.POST("/security/password", controller.ChangePassword)
	}

	return r
}