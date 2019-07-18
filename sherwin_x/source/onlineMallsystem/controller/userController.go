package controller

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"onlineMallsystem/conf/Err"
	"onlineMallsystem/conf/msg"
	"onlineMallsystem/model"
)

//注册:POST
//localhost:8080/user
func SignIn(c *gin.Context) {
	log.Println(">>>User Signing Up<<<")
	newUser := msg.User{}

	//bind sign massage
	if err := c.ShouldBind(&newUser); err != nil {
		c.JSON(200, Err.BindingFailedJson)
		return
	}

	//check id availability
	filter := bson.M{"username": newUser.Username}
	if _, err := model.FindUser(filter); err == nil {
		c.JSON(200, Err.UserExistJson)
		return
	}

	//encode psw to md5 before insert
	pswMd5 := md5.New()
	pswMd5.Write([]byte(newUser.Psw))
	newUser.Psw = string(pswMd5.Sum(nil))

	//insert new user
	if err := model.InsertUser(newUser); err != nil {
		c.JSON(200, Err.InsertFailedJson)
		return
	}
	c.JSON(200, gin.H{
		"success": true,
		"error":   "",
		"data":    "ok"})
}

//登录:POST
//localhost:8080/user/login
//引用jwt库中自定函数

//查看某位用户资料
func ShowUser(c *gin.Context) {
	log.Println(">>>Get User's Message<<<")
	var userId string
	//binding key
	userId = c.Param("id")
	if userId == "" {
		c.JSON(200, Err.NoKeyJson)
		return
	}
	//get my id
	if myId, err := c.Get(msg.IdentityKey); !err {
		log.Println("warning: id get failed")
		return
	} else if userId == myId {
		ShowMe(c)
		return
	}
	//find user
	stringId := fmt.Sprintf("%v", userId)
	ojId, _ := primitive.ObjectIDFromHex(stringId)
	res, err := model.FindUser(bson.M{"_id": ojId})
	if err != nil {
		c.JSON(200, Err.UserNotExistJson)
		return
	}
	//view_count+1
	if err := model.UserUpdate(ojId, "total_view_count", res.TotalViewCount+1); err != nil {
		c.JSON(200, Err.GetFailedJson)
		return
	}
	c.JSON(200, gin.H{
		"success": true,
		"error":   "",
		"data": gin.H{
			"nickname":            res.Nickname,
			"email":               res.Email,
			"total_view_count":    res.TotalViewCount,
			"total_collect_count": res.TotalCollectCount,
		}})

}

//查看个人资料:GET
//localhost:8080/me
func ShowMe(c *gin.Context) {
	log.Println(">>>Get Himself's Message<<<")
	id, err := c.Get(msg.IdentityKey)
	if !err {
		log.Println("warning: id get failed")
		return
	}
	stringId := fmt.Sprintf("%v", id)
	ojId, _ := primitive.ObjectIDFromHex(stringId)
	if res, err := model.FindUser(bson.M{"_id": ojId}); err != nil {
		c.JSON(200, Err.GetFailedJson)
	} else {
		c.JSON(200, gin.H{
			"success": true,
			"error":   "",
			"data": gin.H{
				"username":            res.Username,
				"nickname":            res.Nickname,
				"mobile":              res.Mobile,
				"email":               res.Email,
				"total_view_count":    res.TotalViewCount,
				"total_collect_count": res.TotalCollectCount,
			}})
	}
}

//修改个人资料:POST
//localhost:8080/me
//TBD
func UpdateMe(c *gin.Context) {
	log.Println(">>>User Update Message<<<")
	newData := msg.User{}
	//bind sign massage
	if err := c.ShouldBind(&newData); err != nil {
		c.JSON(200, Err.BindingFailedJson)
		return
	}
	if newData.Username != "" {
	}
}

//查看我的发布:GET
//localhost:8080/me/commodities
func MyCommodities(c *gin.Context) {
	log.Println(">>>Get My Commodities<<<")
	id, err := c.Get(msg.IdentityKey)
	if !err {
		log.Println("warning: id get failed")
	}
	stringId := fmt.Sprintf("%v", id)
	//log.Println("\n", id, "\n", stringId)
	if res, err := model.FindAllCommodity(bson.M{"pub_id": stringId}); err != nil {
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

//查看我的收藏:GET
func MyCollections(c *gin.Context) {
	log.Println(">>>Get My Collection<<<")
	id, err := c.Get(msg.IdentityKey)
	if !err {
		log.Println("warning: id get failed")
	}
	stringId := fmt.Sprintf("%v", id)
	//log.Println("\n", id, "\n", stringId)
	//ojId, _ := primitive.ObjectIDFromHex(stringId)
	if res, err := model.FindAllCollection(bson.M{"user_id": stringId}); err != nil {
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

//收藏某个商品:POST
func NewCollection(c *gin.Context) {
	log.Println(">>>New Collection<<<")
	newCollection := msg.Collection{}
	//bind sign massage
	if err := c.ShouldBind(&newCollection); err != nil {
		c.JSON(200, Err.BindingFailedJson)
		return
	}
	//check commodity exist
	ojId, _ := primitive.ObjectIDFromHex(newCollection.Id)
	res, err := model.FindOneCommodity(bson.M{"_id": ojId})
	if err != nil {
		c.JSON(200, Err.GetFailedJson)
		return
	}
	//get user id
	id, exist := c.Get(msg.IdentityKey)
	if !exist {
		log.Println("warning: id get failed")
	}
	stringId := fmt.Sprintf("%v", id)
	//log.Println(stringId)
	//check if user have collected commodity
	filter := bson.M{"user_id": stringId, "id": newCollection.Id}
	if _, err := model.FindOneCollection(filter); err == nil {
		c.JSON(200, Err.CollectionExistJson)
		return
	}
	//insert
	newCollection.UserId = stringId
	newCollection.Title = res.Title
	if err := model.InsertCollection(newCollection); err != nil {
		c.JSON(200, Err.InsertFailedJson)
		return
	}
	//commodity collect_count+1
	if err := model.CommodityUpdate(ojId, "collect_count", res.CollectCount+1); err != nil {
		c.JSON(200, Err.GetFailedJson)
		return
	}
	//user collect_count+1
	ojUserId, _ := primitive.ObjectIDFromHex(stringId)
	user, _ := model.FindUser(bson.M{"_id": ojUserId})
	if err := model.UserUpdate(ojUserId, "total_collect_count", user.TotalCollectCount+1); err != nil {
		c.JSON(200, Err.GetFailedJson)
		return
	}
	c.JSON(200, gin.H{
		"success": true,
		"error":   "",
		"data":    "ok"})
}

//删除某个收藏:DELETE
func DeleteCollection(c *gin.Context) {
	log.Println(">>>Delete Collection<<<")
	delCollection := msg.Collection{}
	//bind sign massage
	if err := c.ShouldBind(&delCollection); err != nil {
		c.JSON(200, Err.BindingFailedJson)
		return
	}
	//get user id
	id, exist := c.Get(msg.IdentityKey)
	if !exist {
		log.Println("warning: id get failed")
	}
	stringId := fmt.Sprintf("%v", id)
	filter := bson.M{"user_id": stringId, "id": delCollection.Id}
	if err := model.DeleteOneCollection(filter); err != nil {
		c.JSON(200, Err.DeleteFailedJson)
		return
	}
	c.JSON(200, gin.H{
		"success": true,
		"error":   "",
		"data":    "ok",})
}
