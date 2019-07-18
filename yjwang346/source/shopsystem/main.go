package main

import (
	"github.com/gin-gonic/gin"

)

func amain(){
	router :=gin.New()


	router.Run(":8080")
}