package main

import (
	"database/sql"
	//"fmt"
	"github.com/gin-gonic/gin"
	//"project1/database"
	"project1/controllers"
)
var db *sql.DB		//通过查询，知道db是*sql.DB类型

func main(){
	r := gin.New()
	//controllers.InitMysql()
	//user :=r.Group("/user")
	r.POST("/login",controllers.Signup)
	/*
	admin :=r.Group("/admin")
	admin.GET("/getuserinfo",controllers.Getuserinfo)
	admin.GET("/getusersinfo",controllers.GetUsersinfo)
	admin.POST("/deleteuser",controllers.Deletuser)
	 */

}