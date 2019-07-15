package controller

import (
	"awesomeProject/database"
	"awesomeProject/model"
	"context"
	"crypto/md5"
	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

var UserColl *mongo.Collection
var ctx context.Context

//localhost:8080/sign　　
func SignHandler(c *gin.Context) {
	log.Println(">>>User Signing Up<<<")
	newuser := model.SignForm{}

	//bind sign massage
	if err := c.ShouldBind(&newuser); err != nil {
		c.JSON(400, gin.H{"warning": "invalid massage"})
		return
	}

	//check id availability
	filter := bson.M{"id": newuser.Id}
	if _, err := database.FindUser(filter); err == nil {
		c.JSON(400, gin.H{"warning": "invalid id!"})
		return
	}

	//encode psw to md5 before insert
	newuser.Psw = encode(newuser.Psw)

	//insert new user
	if err := database.InsertUser(newuser); err != nil {
		c.JSON(500, gin.H{"statu": "sorry, sign up failed!"})
	}
	c.JSON(200, gin.H{"statu": "sign up success!"})
}

//encode user's password
func encode(psw string) string {
	pswMd5 := md5.New()
	pswMd5.Write([]byte(psw))
	psw = string(pswMd5.Sum(nil))
	return psw
}

//user authority
func UserCallback(c *gin.Context) (interface{}, error) {
	log.Println(">>>User Authoring<<<")
	user := model.LoginForm{}

	//bind login massage
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(400, gin.H{"warning": "invalid massage"})
		return nil, nil
	}

	//encode psw to md5
	user.Psw = encode(user.Psw)

	//find user in db
	filter := bson.M{
		"id":  user.Id,
		"psw": user.Psw}
	if _, err := database.FindUser(filter); err != nil {
		return nil, jwt.ErrFailedAuthentication
	}
	return &model.LoginForm{Id: user.Id,}, nil
}

//localhost:8080/user/hello
func HelloUserHandler(c *gin.Context) {
	log.Println(">>>User Auth Test<<<")
	id, err := c.Get(model.IdentityKey)
	if !err {
		log.Println("warning: id get failed")
	}
	c.JSON(200, gin.H{
		"user": id,
		"text": "Welcome!",
	})
}

//localhost:8080/user/update　　　
func UpdateHandler(c *gin.Context) {
	log.Println(">>>User Message Update<<<")
	id, err := c.Get(model.IdentityKey)
	if !err {
		log.Println("warning: id get failed")
	}
	newdate := model.UpdateForm{}

	//check massage availability
	if err := c.ShouldBind(&newdate); err != nil {
		c.JSON(400, gin.H{"warning": "invalid massage"})
		return
	}
	if newdate.Item != "psw" && newdate.Item != "name" && newdate.Item != "tel" && newdate.Item != "email" {
		c.JSON(400, gin.H{"warning": "invalid item"})
		return
	}

	//update massage
	if err := database.UpdateMsg(id, newdate); err != nil {
		c.JSON(200, gin.H{"state": "massage not change"})
		return
	}
	c.JSON(200, gin.H{"state": "update success!"})
}
