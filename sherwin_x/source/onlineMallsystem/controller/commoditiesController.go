package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"onlineMallsystem/conf/Err"
	"onlineMallsystem/conf/msg"
	"onlineMallsystem/model"
)

//获取商品列表
func GetCommodities(c *gin.Context) {
	log.Println(">>>Get Commodity List<<<")
	list := msg.GetCommodity{}
	//bind massage
	if err := c.ShouldBind(&list); err != nil {
		c.JSON(200, Err.BindingFailedJson)
		return
	}
	var filter bson.M
	switch {
	case list.Category == 0 && list.Keyword == "":
		filter = bson.M{"": nil}
	case list.Category != 0 && list.Keyword == "":
		filter = bson.M{"category": list.Category}
	case list.Category == 0 && list.Keyword != "":
		filter = bson.M{"title": primitive.Regex{Pattern: list.Keyword, Options: ""}}
	case list.Category != 0 && list.Keyword != "":
		filter = bson.M{"category": list.Category, "title": primitive.Regex{Pattern: list.Keyword, Options: ""}}
	}
	if res, err := model.GetCommoditiesList(list.Page, list.Limit, filter); err != nil {
		c.JSON(200, gin.H{
			"success": true,
			"error":   "",
			"data":    ""})
	} else {
		//data:=map[string]interface{}{"id":res.id,"title": res}
		c.JSON(200, gin.H{
			"success": true,
			"error":   "",
			"data":    res})
	}
}

//获取热搜词
func GetHotSearch(c *gin.Context) {

}

//发布新商品
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
		log.Println("warning: id get failed")
	}
	newCommodity.PubUser = fmt.Sprintf("%v", id)

	//insert commodity
	if err := model.InsertCommodity(newCommodity); err != nil {
		c.JSON(200, Err.InsertFailedJson)
		return
	}
	c.JSON(200, gin.H{
		"success": true,
		"error":   "",
		"data":    "ok"})
}

//某个商品详情 GET
//localhost:8080/commodity/:id
func DetailCommodity(c *gin.Context) {
	log.Println(">>>Get Commodity Detail<<<")
	var id string
	//binding
	id = c.Param("id")
	if id == "" {
		c.JSON(200, Err.NoKeyJson)
		return
	}
	ojId, _ := primitive.ObjectIDFromHex(id)
	res, err := model.FindOneCommodity(bson.M{"_id": ojId})
	if err != nil {
		c.JSON(200, Err.GetFailedJson)
		return
	}
	//view_count+1
	if err := model.CommodityUpdate(ojId, "view_count", res.ViewCount+1); err != nil {
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
			"view_count":    res.ViewCount,
			"collect_count": res.CollectCount,
		}})

}

//删除某个商品
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
		log.Println("warning: id get failed")
		return
	}
	ojId, _ := primitive.ObjectIDFromHex(commodityId)
	if _, err := model.FindOneCommodity(bson.M{"_id": ojId}); err != nil {
		c.JSON(200, Err.CommodityNotExistJson)
		return
	}
	if _, err := model.FindOneCommodity(bson.M{"_id": ojId, "pub_id": myId}); err != nil {
		c.JSON(200, Err.DeleteFailedJson)
		return
	}
	if err := model.DeleteCommodity(bson.M{"_id": ojId, "pub_id": myId}); err != nil {
		c.JSON(200, Err.DeleteFailedJson)
		return
	}
	c.JSON(200, gin.H{
		"success": true,
		"error":   "",
		"data":    "ok",})
}
