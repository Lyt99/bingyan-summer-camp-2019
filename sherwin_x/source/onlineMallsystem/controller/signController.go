package controller

import (
	"crypto/md5"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"onlineMallsystem/conf"
	"onlineMallsystem/model"
)

//localhost:8080/sign　　
func SignHandler(c *gin.Context) {
	log.Println(">>>User Signing Up<<<")
	newuser := conf.User{}

	//bind sign massage
	if err := c.ShouldBind(&newuser); err != nil {
		c.JSON(400, gin.H{"warning": "invalid massage"})
		return
	}

	//check id availability
	filter := bson.M{"tel": newuser.Tel}
	if _, err := model.FindUser(filter); err == nil {
		c.JSON(400, gin.H{"warning": "invalid tel!"})
		return
	}

	//encode psw to md5 before insert
	pswMd5 := md5.New()
	pswMd5.Write([]byte(newuser.Psw))
	newuser.Psw = string(pswMd5.Sum(nil))

	//insert new user
	if err := model.InsertUser(newuser); err != nil {
		c.JSON(500, gin.H{"statu": "sorry, sign up failed!"})
	}
	c.JSON(200, gin.H{"statu": "sign up success!"})
}


