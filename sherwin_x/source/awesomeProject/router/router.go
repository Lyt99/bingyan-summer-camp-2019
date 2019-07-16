package router

import (
	"awesomeProject/controller"
	"awesomeProject/middleware"
	"github.com/gin-gonic/gin"
)

//Demo System Routers
func Router() {
	r := gin.New()
	r.POST("/sign", controller.SignHandler)

	//user router
	user := r.Group("/user")
	user.POST("/login", middleware.GetUserToken().LoginHandler)
	user.Use(middleware.GetUserToken().MiddlewareFunc())
	{
		user.POST("/hello", controller.HelloUserHandler)
		user.POST("/update", controller.UpdateHandler)
	}

	//admin router
	admin := r.Group("/admin")
	admin.POST("/login", middleware.GetAdminToken().LoginHandler)
	admin.Use(middleware.GetAdminToken().MiddlewareFunc())
	{
		admin.POST("/hello", controller.HelloAdminHandler)
		admin.POST("/find", controller.FindHandler)
		admin.POST("/show", controller.ShowHandler)
		admin.POST("/delete", controller.DelHandler)
	}

	_ = r.Run(":8080")
}
