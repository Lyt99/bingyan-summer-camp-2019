package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"signin_Project/database"
)
var db *sql.DB		//通过查询，知道db是*sql.DB类型
func main(){
	r := gin.New()
	database.InitMysql()
	login :=r.Group("/login")
	//signup :=r.Group("/signup")

	r.POST("/signup",)
}