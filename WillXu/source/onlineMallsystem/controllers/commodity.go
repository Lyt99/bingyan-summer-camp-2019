package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"onlineMallsystem/models"
	"onlineMallsystem/models/Err"
	"onlineMallsystem/models/msg"
	"strconv"
)

//获取商品列表:GET
//localhost:8080/commodities
func GetCommodities(c *gin.Context) {
	log.Println(">>>Get Commodity List<<<")
	list := msg.GetCommodity{}
	//bind massage
	list.Page, _ = strconv.Atoi(c.Query("page"))
	list.Limit, _ = strconv.Atoi(c.Query("limit"))
	list.Category, _ = strconv.Atoi(c.DefaultQuery("category", "0"))
	list.Keyword = c.DefaultQuery("keyword", "")
	//keyword
	if list.Keyword != "" {
		key := msg.Key{Keyword: list.Keyword, Count: 1}
		if err := models.KeyFunc(key); err != nil {
			c.JSON(200, Err.GetFailedJson)
			return
		}
	}
	//check input err
	if list.Page <= 0 || list.Limit <= 0||list.Category<0||list.Category>9 {
		c.JSON(200, Err.BindingFailedJson)
		return
	}
	//get list
	var filter bson.M
	switch {
	case list.Category == 0 && list.Keyword == "":
		filter = bson.M{"type": "commodity"}
	case list.Category != 0 && list.Keyword == "":
		filter = bson.M{"category": list.Category}
	case list.Category == 0 && list.Keyword != "":
		filter = bson.M{"title": primitive.Regex{Pattern: list.Keyword, Options: ""}}
	case list.Category != 0 && list.Keyword != "":
		filter = bson.M{"category": list.Category, "title": primitive.Regex{Pattern: list.Keyword, Options: ""}}
	}
	if res, err := models.GetCommoditiesList(list.Page, list.Limit, filter); err != nil {
		c.JSON(200, gin.H{
			"success": true,
			"error":   "",
			"data":    ""})
	} else {
		c.JSON(200, gin.H{
			"success": true,
			"error":   "",
			"data":    res})
	}
}

//获取热搜词:GET
//localhost:8080/commodities/hot
func GetHotSearch(c *gin.Context) {
	if res, err := models.FindAllKeyword(); err != nil {
		log.Println(err)
		c.JSON(200, gin.H{
			"success": true,
			"error":   "",
			"data":    ""})
	} else {
		c.JSON(200, gin.H{
			"success": true,
			"error":   "",
			"data":    res})
	}
}

//发布新商品:POST
//localhost:8080/commodities
func NewCommodity(c *gin.Context) {
	log.Println(">>>Public New Commodity<<<")
	newCommodity := msg.Commodity{}
	//bind massage
	if err := c.ShouldBind(&newCommodity); err != nil {
		c.JSON(200, Err.BindingFailedJson)
		return
	}
	//write pub_user id
	id, err := c.Get(msg.IdentityKey)
	if !err {
		c.JSON(200, Err.IdGetFailedJson)
		return
	}
	newCommodity.PubUser = fmt.Sprintf("%v", id)
	//insert commodity
	if err := models.InsertCommodity(newCommodity); err != nil {
		c.JSON(200, Err.InsertFailedJson)
		return
	}
	c.JSON(200, successJson)
}

//某个商品详情 GET
//localhost:8080/commodity/:id
func DetailCommodity(c *gin.Context) {
	log.Println(">>>Get Commodity Detail<<<")
	var id string
	//binding
	id = c.Param("id")
	//check commodity exist
	ojId, _ := primitive.ObjectIDFromHex(id)
	res, err := models.FindOneCommodity(bson.M{"_id": ojId})
	if err != nil {
		c.JSON(200, Err.CommodityNotExistJson)
		return
	}
	//view_count+1
	if err := models.CommodityUpdate(ojId, "view_count", res.ViewCount+1); err != nil {
		c.JSON(200, Err.GetFailedJson)
		return
	}
	c.JSON(200, gin.H{
		"success": true,
		"error":   "",
		"data": gin.H{
			"pub_user":      res.PubUser,
			"title":         res.Title,
			"desc":          res.Desc,
			"category":      res.Category,
			"price":         res.Price,
			"picture":       res.Picture,
			"view_count":    res.ViewCount + 1,
			"collect_count": res.CollectCount,
		}})

}

//删除某个商品:DELETE
//localhost:8080/commodity/:id
func DeleteCommodity(c *gin.Context) {
	log.Println(">>>Delete User's Commodity<<<")
	var commodityId string
	//binding key
	commodityId = c.Param("id")
	if commodityId == "" {
		c.JSON(200, Err.NoKeyJson)
		return
	}
	//get my id
	myId, err := c.Get(msg.IdentityKey)
	if !err {
		c.JSON(200, Err.IdGetFailedJson)
		return
	}
	//check commodity exist
	ojId, _ := primitive.ObjectIDFromHex(commodityId)
	if _, err := models.FindOneCommodity(bson.M{"_id": ojId}); err != nil {
		c.JSON(200, Err.CommodityNotExistJson)
		return
	}
	//check delete authority
	if _, err := models.FindOneCommodity(bson.M{"_id": ojId, "pub_id": myId}); err != nil {
		c.JSON(200, Err.DeleteFailedJson)
		return
	}
	if err := models.DeleteCommodity(bson.M{"_id": ojId, "pub_id": myId}); err != nil {
		c.JSON(200, Err.GetFailedJson)
		return
	}
	c.JSON(200, successJson)
}
