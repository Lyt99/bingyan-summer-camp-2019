package main

import (
	"./controller"
	"github.com/gin-gonic/gin"
)


func main() {
	r := gin.Default()
	authMiddleware := controller.MiddleWareInit() // this middle ware cannot tell the authority
													// authority is evaluated by the handlers!

	r.GET("/", controller.MainPage)
	r.POST("/register", controller.Register)
	r.POST("/login", authMiddleware.LoginHandler)

	admin := r.Group("/admin")
	r.POST("/register/admin", controller.AdminRegister)
	admin.Use(authMiddleware.MiddlewareFunc())
	{
		admin.GET("/manage/:pageno", controller.AdminManage)
	}

	r.GET("/user/:username", controller.InfoPage)
	user := r.Group("/user")
	user.Use(authMiddleware.MiddlewareFunc())
	{
		user.POST("/center/:username", controller.UpdateHandler)
		user.POST("/security/password", controller.ChangePassword)
	}

	r.Run(":8080")
}