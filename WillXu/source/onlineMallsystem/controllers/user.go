package controllers

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"onlineMallsystem/models"
	"onlineMallsystem/models/Err"
	"onlineMallsystem/models/msg"
)

var successJson = map[string]interface{}{
	"success": true,
	"error":   "",
	"data":    "ok"}

//>>>>>注册与登录接口无中间件认证<<<<<
//注册:POST
//localhost:8080/user
func SignUp(c *gin.Context) {
	log.Println(">>>Sign Up<<<")
	newUser := msg.User{}
	//bind sign massage
	if err := c.ShouldBind(&newUser); err != nil {
		c.JSON(200, Err.BindingFailedJson)
		return
	}
	//check username availability
	filter := bson.M{"username": newUser.Username}
	if _, err := models.FindUser(filter); err == nil {
		c.JSON(200, Err.UserExistJson)
		return
	}
	//encode psw to md5 before insert
	pswMd5 := md5.New()
	pswMd5.Write([]byte(newUser.Psw))
	newUser.Psw = string(pswMd5.Sum(nil))
	//insert new user
	if err := models.InsertUser(newUser); err != nil {
		c.JSON(200, Err.InsertFailedJson)
		return
	}
	c.JSON(200, successJson)
}

//登录:POST
//localhost:8080/user/login
//引用jwt库中自定函数LoginHandler

//>>>>>下方接口均添加jwt中间件认证<<<<<
//查看某位用户资料:Get
//localhost:8080/user/:id
func ShowUser(c *gin.Context) {
	log.Println(">>>Get User Message<<<")
	var userId string
	//binding :id
	userId = c.Param("id")
	//get my id
	if myId, err := c.Get(msg.IdentityKey); !err {
		c.JSON(200, Err.IdGetFailedJson)
		return
	} else if userId == myId {
		ShowMe(c)
		return
	}
	stringId := fmt.Sprintf("%v", userId)
	ojId, _ := primitive.ObjectIDFromHex(stringId)
	//find user
	res, err := models.FindUser(bson.M{"_id": ojId})
	if err != nil {
		c.JSON(200, Err.UserNotExistJson)
		return
	}
	//total_view_count+1
	if err := models.UserUpdate(ojId, "total_view_count", res.TotalViewCount+1); err != nil {
		c.JSON(200, Err.GetFailedJson)
		return
	}
	c.JSON(200, gin.H{
		"success": true,
		"error":   "",
		"data": gin.H{
			"nickname":            res.Nickname,
			"email":               res.Email,
			"total_view_count":    res.TotalViewCount + 1,
			"total_collect_count": res.TotalCollectCount,
		}})

}

//查看个人资料:GET
//localhost:8080/me
func ShowMe(c *gin.Context) {
	log.Println(">>>Get My Message<<<")
	//get user id
	id, err := c.Get(msg.IdentityKey)
	if !err {
		c.JSON(200, Err.IdGetFailedJson)
		return
	}
	stringId := fmt.Sprintf("%v", id)
	ojId, _ := primitive.ObjectIDFromHex(stringId)
	//find user in db
	if res, err := models.FindUser(bson.M{"_id": ojId}); err != nil {
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
func UpdateMe(c *gin.Context) {
	log.Println(">>>Update My Message<<<")
	newData := msg.User{}
	//bind sign massage
	if err := c.ShouldBind(&newData); err != nil {
		c.JSON(200, Err.BindingFailedJson)
		return
	}
	//get user id
	id, exist := c.Get(msg.IdentityKey)
	if !exist {
		c.JSON(200, Err.IdGetFailedJson)
	}
	stringId := fmt.Sprintf("%v", id)
	ojId, _ := primitive.ObjectIDFromHex(stringId)
	//find user
	res, err := models.FindUser(bson.M{"_id": ojId})
	if err != nil {
		c.JSON(200, Err.GetFailedJson)
		return
	}
	//update
	if newData.Username != res.Username {
		//check username availability
		filter := bson.M{"username": newData.Username}
		if _, err := models.FindUser(filter); err == nil {
			c.JSON(200, Err.UserExistJson)
			return
		}
		if err := models.UpdateMsg(ojId, "username", newData.Username); err != nil {
			c.JSON(200, Err.GetFailedJson)
			return
		}
	}
	if newData.Nickname != res.Nickname {
		if err := models.UpdateMsg(ojId, "nickname", newData.Nickname); err != nil {
			c.JSON(200, Err.GetFailedJson)
			return
		}
	}
	if newData.Mobile != res.Mobile {
		if err := models.UpdateMsg(ojId, "mobile", newData.Mobile); err != nil {
			c.JSON(200, Err.GetFailedJson)
			return
		}
	}
	if newData.Email != res.Email {
		if err := models.UpdateMsg(ojId, "email", newData.Email); err != nil {
			c.JSON(200, Err.GetFailedJson)
			return
		}
	}
	if newData.Psw != res.Psw {
		if err := models.UpdateMsg(ojId, "psw", newData.Psw); err != nil {
			c.JSON(200, Err.GetFailedJson)
			return
		}
	}
	c.JSON(200, successJson)
}

//查看我的发布:GET
//localhost:8080/me/commodities
//TO-DO:返回键开头大写
func MyCommodities(c *gin.Context) {
	log.Println(">>>Get My Commodities<<<")
	//get user id
	id, err := c.Get(msg.IdentityKey)
	if !err {
		c.JSON(200, Err.IdGetFailedJson)
		return
	}
	stringId := fmt.Sprintf("%v", id)
	if res, err := models.FindAllCommodity(bson.M{"pub_id": stringId}); err != nil {
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

//查看我的收藏:GET
//localhost:8080/me/collections
//TO-DO:返回键开头大写
func MyCollections(c *gin.Context) {
	log.Println(">>>Get My Collection<<<")
	id, err := c.Get(msg.IdentityKey)
	if !err {
		c.JSON(200, Err.IdGetFailedJson)
	}
	stringId := fmt.Sprintf("%v", id)
	if res, err := models.FindAllCollection(bson.M{"user_id": stringId}); err != nil {
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

//收藏某个商品:POST
//localhost:8080/me/collections
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
	res, err := models.FindOneCommodity(bson.M{"_id": ojId})
	if err != nil {
		c.JSON(200, Err.CommodityNotExistJson)
		return
	}
	//get user id
	id, exist := c.Get(msg.IdentityKey)
	if !exist {
		c.JSON(200, Err.IdGetFailedJson)
	}
	stringId := fmt.Sprintf("%v", id)
	//check if user have collected commodity
	filter := bson.M{"user_id": stringId, "id": newCollection.Id}
	if _, err := models.FindOneCollection(filter); err == nil {
		c.JSON(200, Err.CollectionExistJson)
		return
	}
	//insert
	newCollection.UserId = stringId
	newCollection.Title = res.Title
	if err := models.InsertCollection(newCollection); err != nil {
		c.JSON(200, Err.InsertFailedJson)
		return
	}
	//commodity collect_count+1
	if err := models.CommodityUpdate(ojId, "collect_count", res.CollectCount+1); err != nil {
		c.JSON(200, Err.GetFailedJson)
		return
	}
	//user collect_count+1
	ojUserId, _ := primitive.ObjectIDFromHex(stringId)
	user, _ := models.FindUser(bson.M{"_id": ojUserId})
	if err := models.UserUpdate(ojUserId, "total_collect_count", user.TotalCollectCount+1); err != nil {
		c.JSON(200, Err.GetFailedJson)
		return
	}
	c.JSON(200, successJson)
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
	//check commodity exist
	ojId, _ := primitive.ObjectIDFromHex(delCollection.Id)
	res, err := models.FindOneCommodity(bson.M{"_id": ojId})
	if err != nil {
		c.JSON(200, Err.CommodityNotExistJson)
		return
	}
	//get user id
	id, exist := c.Get(msg.IdentityKey)
	if !exist {
		c.JSON(200, Err.IdGetFailedJson)
		return
	}
	stringId := fmt.Sprintf("%v", id)
	filter := bson.M{"user_id": stringId, "id": delCollection.Id}
	if err := models.DeleteOneCollection(filter); err != nil {
		c.JSON(200, Err.DeleteFailedJson)
		return
	}
	//commodity collect_count+1
	if err := models.CommodityUpdate(ojId, "collect_count", res.CollectCount-1); err != nil {
		c.JSON(200, Err.GetFailedJson)
		return
	}
	//user collect_count+1
	ojUserId, _ := primitive.ObjectIDFromHex(stringId)
	user, _ := models.FindUser(bson.M{"_id": ojUserId})
	if err := models.UserUpdate(ojUserId, "total_collect_count", user.TotalCollectCount-1); err != nil {
		c.JSON(200, Err.GetFailedJson)
		return
	}
	c.JSON(200, successJson)
}
