package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shopsystem/database"
)

//7-23,并没有进行debug
func Get_hot_keyword(c *gin.Context)  {
	var keywords []string
	keywords,i := database.GetHotkeyword(c)
	if i == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"success":false,
			"error": "查询热门关键词出现错误",
			"data":"",//失败的时候留空
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"success":true,
		"error": "",
		"data":keywords,//失败的时候留空
	})
}
