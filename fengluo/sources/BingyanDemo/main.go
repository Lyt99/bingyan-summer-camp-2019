package main

import (
	"BingyanDemo/database"
	"BingyanDemo/routers"
)

func main()  {
	database.InitMysql()
	routers :=routers.InitRouter()

	routers.Run(":8080")
}