package main

import (
	"github.com/gin-gonic/gin"
	"shopsystem/controller"
)

func main(){
	router :=gin.New()

	router.POST("/register",controller.Signup)

	router.Run(":8080")
}