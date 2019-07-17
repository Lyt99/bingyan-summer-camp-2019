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
	myId, err := c.Get(msg.IdentityKey)
	if !err {
		log.Println("warning: id get failed")
		return
	}
	//if key=myId, jump to ShowMeFunc
	if userId==myId{
		ShowMe(c)
		return
	}
	//find user
	stringId := fmt.Sprintf("%v", userId)
	ojId, _ := primitive.ObjectIDFromHex(stringId)
	res, err := model.FindUser(bson.M{"_id": ojId})
	if err != nil {
		c.JSON(200, Err.UserNotExistJson)
	} else {
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

}

//查看我的发布:GET
//localhost:8080/me/commodities
//TBD
func MyCommodities(c *gin.Context) {
	log.Println(">>>Get My Commodities<<<")
	id, err := c.Get(msg.IdentityKey)
	if !err {
		log.Println("warning: id get failed")
	}
	stringId := fmt.Sprintf("%v", id)
	log.Println("\n", id, "\n", stringId)
	if res, err := model.FindAllCommodity(bson.M{"pub_id": stringId}); err != nil {
		c.JSON(200, Err.GetFailedJson)
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

}

//收藏某个商品:POST
func NewCollection(c *gin.Context) {

}

//删除某个收藏:DELETE
func DeleteCollection(c *gin.Context) {

}
