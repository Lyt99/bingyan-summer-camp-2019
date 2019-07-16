package main

import (
	"database/sql"
	"project1/middleware"

	//"fmt"
	"github.com/gin-gonic/gin"
	//"project1/database"
	"project1/controllers"
)
//var db *sql.DB		//通过查询，知道db是*sql.DB类型

func main(){
	r := gin.New()
	//controllers.InitMysql()
	user :=r.Group("/user")
	r.POST("/login",controllers.Signup)

	user.POST("/changeinfomation",controllers.Changeinfo)

	admin :=r.Group("/admin")
	admin.Use(middleware.AdminMiddleWareInit())
	admin.GET("/getuserinfo",controllers.Getuserinfo)
	admin.GET("/getusersinfo",controllers.GetUsersinfo)
	admin.POST("/deleteuser",controllers.Deletuser)


}