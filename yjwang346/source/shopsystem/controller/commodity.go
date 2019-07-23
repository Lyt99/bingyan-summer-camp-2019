package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shopsystem/database"
	"shopsystem/model"
	"strconv"
)

func Postcommodity(c *gin.Context){
	//从token处得到用户名
	id,errr := c.Get("id")
	username := id.(string)
	if !errr {
		c.JSON(http.StatusBadRequest, gin.H{
			"success":false,
			"error": "获得用户id出现错误",
			"data":"",//失败的时候留空
		})
		return
	}
	var commodity model.PostCommodity
	//获得相关的信息
	if err := c.ShouldBindJSON(&commodity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success":false,
			"error": err.Error(),
			"data":"",//失败的时候留空
		})
		return
	}
	//加入数据库中,相关的检测之类的可以后面再加
	database.CreateCommodity(c,username,commodity.Title,commodity.Description,commodity.Picture,commodity.Price,commodity.Category)

}

//删除商品通过id来进行删除
func Deletecommodity(c *gin.Context){
	deleteid := c.Param("id")
	id,err := strconv.Atoi(deleteid)
	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"success":false,
			"error": "删除商品失败,传递id失败",
			"data":"",//失败的时候留空
		})
		return
	}
	database.DeleteCommo(c,id)
}

//查询商品的信息，先放一放把，查询方面的事之后再说
func  Getcommodityinfo(c *gin.Context){

}