package main

import (
	"github.com/gin-gonic/gin"
)

func main(){
	r:=gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/signpost", func(c *gin.Context) {
		c.HTML(200,"postForm.html",gin.H{})
	})
	//r.POST("/post",RenderForm)
	_ = r.Run(":8080")
}
