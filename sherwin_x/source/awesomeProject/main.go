package main

import (
	"awesomeProject/controller"
	"awesomeProject/middleware"
	"github.com/gin-gonic/gin"
)

//To do
//1.
func main(){
	r:=gin.New()

	//Router
	r.POST("/sign",controller.SignHandler)

	user:=r.Group("/user")
	userToken:=middleware.GetUserToken()
	user.POST("/login",userToken.LoginHandler)
	user.Use(userToken.MiddlewareFunc())
	{
		user.POST("/hello",controller.HelloUserHandler)
		user.POST("/update",controller.UpdateHandler)
	}

	admin:=r.Group("/admin")
	adminToken:=middleware.GetAdminToken()
	admin.POST("/login",adminToken.LoginHandler)
	admin.Use(adminToken.MiddlewareFunc())
	{
		admin.GET("/hello",controller.HelloAdminHandler)
		admin.POST("/find",controller.FindHandler)
		admin.POST("/show",controller.ShowHandler)
		admin.POST("/delete",controller.DelHandler)
	}

	_ = r.Run(":8080")
}