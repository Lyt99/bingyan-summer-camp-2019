package main

import (
	"awesomeProject/controller"
	"github.com/gin-gonic/gin"
)

//测试注册网页
//>>>>path need update<<<<
func main() {
	r := gin.Default()
	r.LoadHTMLGlob("test/templates/*")
	r.GET("/sign", func(c *gin.Context) {
		c.HTML(200, "postForm.html", gin.H{})
	})
	r.POST("/sign", controller.SignHandler)
	_ = r.Run(":8080")
}
