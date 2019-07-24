package controller

import (
	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"onlineShopping/model"
	"strconv"
)

//this func is used to get all the
//current user's collections
func GetMyCollections(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	identity, _ := claims["id"]
	userID, _ := identity.(float64)
	response := ResponseInit()
	collection, err := model.DBSearchCollections(int(userID))
	if err != nil {
		dbError(c)
		return
	}
	response["success"] = true
	response["error"] = ""
	response["data"] = collection
	c.JSON(200, response)
}

func AddToCollections(c *gin.Context) {
	var json model.CollectionJSON
	response := ResponseInit()
	claims := jwt.ExtractClaims(c)
	identity, _ := claims["id"].(float64)
	userID := int(identity)
	if err := c.ShouldBindJSON(&json); err != nil {
		response["error"] = "商品ID输入有误"
		c.JSON(400, response)
		return
	}
	commodityID, err := strconv.Atoi(json.CommodityID)
	if err != nil {
		response["error"] = "商品ID输入有误"
		c.JSON(400, response)
		return
	}

	pubID, err := model.DBIsCommodityExist(commodityID)
	if err != nil {
		dbError(c)
		return
	}
	if pubID == 0 {
		response["error"] = "不存在该商品"
		c.JSON(400, response)
		return
	}

	exist, err := model.DBIsCollectionExist(userID, commodityID)
	if err != nil {
		dbError(c)
		return
	}
	if !exist {
		err = model.DBCollectCommodity(userID, commodityID)
		if err != nil {
			dbError(c)
			return
		}
	}
	response["success"] = true
	response["error"] = ""
	response["data"] = "ok"
	c.JSON(200, response)
}

func DeleteFromCollections(c *gin.Context) {
	var json model.CollectionJSON
	response := ResponseInit()
	claims := jwt.ExtractClaims(c)
	identity, _ := claims["id"].(float64)
	userID := int(identity)
	if err := c.ShouldBindJSON(&json); err != nil {
		response["error"] = "商品ID输入有误"
		c.JSON(400, response)
		return
	}
	commodityID, err := strconv.Atoi(json.CommodityID)
	if err != nil {
		response["error"] = "商品ID输入有误"
		c.JSON(400, response)
		return
	}
	exist, err := model.DBIsCollectionExist(userID, commodityID)
	if err != nil {
		dbError(c)
		return
	}
	if !exist {
		response["error"] = "目标商品未收藏"
		c.JSON(400, response)
		return
	}
	if err := model.DBDeleteFromCollections(userID, commodityID); err != nil {
		dbError(c)
		return
	}
	response["success"] = true
	response["error"] = ""
	response["data"] = "ok"
	c.JSON(200, response)
}
