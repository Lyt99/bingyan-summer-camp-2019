package main

import (
	"awesomeProject/controller"
	"awesomeProject/middleware"
	"github.com/gin-gonic/gin"
)

/*成员管理系统
实现内容：
- 管理员和普通用户
- 用户注册和登录
- 用户信息包括用户ID、密码（数据库中加密）、昵称、手机号、邮箱地址

- 管理员
删除普通用户
获取一个成员、所有成员信息

- 普通用户
更改个人信息
 */
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