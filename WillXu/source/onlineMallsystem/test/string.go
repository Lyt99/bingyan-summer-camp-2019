package main

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
	"strings"
)

func main() {
	r := gin.Default()
	r.GET("/get", GetCommodities)   //获取商品列表
	_ = r.Run(":8080")
}


//获取商品列表:GET
//localhost:8080/commodities
func GetCommodities(c *gin.Context) {
	log.Println(">>>Get Commodity List<<<")
	list := msg.GetCommodity{}
	//bind massage
	list.Page, _ = strconv.Atoi(c.Query("page"))
	list.Limit, _ = strconv.Atoi(c.Query("limit"))
	list.Category, _ = strconv.Atoi(c.DefaultQuery("category", "0"))
	list.Keyword = strings.TrimSpace(c.DefaultQuery("keyword", ""))
	//keyword
	if list.Keyword != "" {
		key := msg.Key{Keyword: list.Keyword, Count: 1}
		if err := models.KeyFunc(key); err != nil {
			c.JSON(200, Err.GetFailedJson)
			return
		}
	}
	//check input err
	if list.Page < 0 || list.Limit <= 0 || list.Category < 0 || list.Category > 9 {
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
		filter = bson.M{"title": primitive.Regex{Pattern: list.Keyword, Options: ""}}d
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

func Slece(key string)  {
	//根据空白符分割,不限定中间间隔几个空白符
	fieldsSlece := strings.Fields(key)
	fmt.Println(fieldsSlece)
	for _,v:=range fieldsSlece{
		println(v)
	}
}
