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
	//username := c.Query("username")
	var commodity model.PostCommodity
	//获得相关的信息
	if err := c.ShouldBindJSON(&commodity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success":false,
			"error": err.Error(),//"",
			"data":"传递request错误0",//失败的时候留空
		})
		return
	}
	//加入数据库中,相关的检测之类的可以后面再加
	database.CreateCommodity(c,username,commodity.Title,commodity.Description,commodity.Picture,commodity.Price,commodity.Category)
}

//删除商品通过id来进行删除
func Deletecommodity(c *gin.Context){
	deleteid := c.Query("id")
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
func  Getone_commodityinfo(c *gin.Context){
	uid := c.Query("id")
	id,err := strconv.Atoi(uid)
	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"success":false,
			"error": "获得商品信息失败",
			"data":"",//失败的时候留空
		})
		return
	}
	database.Getone_Commodity(c,id)
}

func Getcommodities(c *gin.Context){
	var findrequest model.Findcommodity
	//很奇怪，命名数据传进来去了，但是还是返回 err 不为 nil
	if err := c.ShouldBindJSON(&findrequest); err != nil && findrequest.Limit == 0{
		c.JSON(http.StatusBadRequest, gin.H{
			"success":false,
			"error": err.Error(),//"",
			"data":"传递搜索请求错误0",//失败的时候留空
		})
		return
	}

	commodities,err := database.Commodities(c,findrequest)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"success":false,
			"error": err,//"",
			"data":"获得商品列表请求错误0",//失败的时候留空
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"success":true,
		"error": "",//"",
		"data":commodities,//这里不是很符合API的格式，但是要传输的是一个列表
		//其他的传输方式似乎还没有这个方法好
	})
	//将搜索的关键词存入数据库中
	if findrequest.Keyword!=""{
		database.Storekeyword(c,findrequest.Keyword)
	}
}