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

//localhost:8080/user
func SignHandler(c *gin.Context) {
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
		"error":"",
		"data":"ok"})
}

//localhost:8080/me
func ShowHandler(c *gin.Context)  {
	log.Println(">>>Get User Message<<<")
	id, err := c.Get(msg.IdentityKey)
	if !err {
		log.Println("warning: id get failed")
	}
	strid:=fmt.Sprintf("%v",id)
	ojid , _ :=primitive.ObjectIDFromHex(strid)
	if res, err := model.FindUser(bson.M{"_id": ojid}); err != nil {
		c.JSON(200, Err.GetFailedJson)
	} else {
		c.JSON(200, gin.H{
			"success":true,
			"error":"",
			"data": gin.H{
				"username":res.Username,
				"nickname":res.Nickname,
				"mobile": res.Mobile,
				"email": res.Email,
				"total_view_count": res.TotalViewCount,
				"total_collect_count": res.TotalCollectCount,
			}})
	}
}

//localhost:8080/me
func UpdateHandler(c *gin.Context)  {
	log.Println(">>>User Update Message<<<")
	newData := msg.User{}

	//bind sign massage
	if err := c.ShouldBind(&newData); err != nil {
		c.JSON(200, Err.BindingFailedJson)
		return
	}

}