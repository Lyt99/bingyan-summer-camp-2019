package controller

import (
	"github.com/gin-gonic/gin"
)

func ResponseInit() Response {
	response := make(Response)
	response["success"] = false
	response["data"] = nil
	return response
}

func internalError(c *gin.Context) {
	response := ResponseInit()
	response["error"] = "内部错误"
	c.JSON(500, response)
}

func dbError(c *gin.Context) {
	response := ResponseInit()
	response["error"] = "无法连接数据库"
	c.JSON(500, response)
}
