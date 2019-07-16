package controller

import (
	"awesomeProject/database"
	"awesomeProject/model"
	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"strconv"
)

//admin authority
func AdminCallback(c *gin.Context) (interface{}, error) {
	log.Println(">>>Admin Authoring<<<")
	user := model.LoginForm{}
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(400, gin.H{"warning": "invalid massage"})
		return nil, nil
	}
	if user.Id == 2019 && user.Psw == "0711" {
		return "200", nil
	}
	return nil, jwt.ErrFailedAuthentication
}

//auth test
//localhost:8080/admin/hello
func HelloAdminHandler(c *gin.Context) {
	log.Println(">>>Admin Auth Test<<<")
	c.JSON(200, gin.H{
		"ID":   "admin",
		"text": "Do What You Want!",
	})
}

//localhost:8080/admin/find
func FindHandler(c *gin.Context) {
	log.Println(">>>Admin Find User<<<")
	findId, _ := strconv.Atoi(c.PostForm("id"))
	filter := bson.M{"id": findId}
	if res, err := database.FindUser(filter); err != nil {
		c.JSON(200, gin.H{"state": "not found"})
	} else {
		c.JSON(200, gin.H{"user": res})
	}
}

//localhost:8080/admin/show
func ShowHandler(c *gin.Context) {
	log.Println(">>>Admin Show All Users<<<")
	if res, err := database.ShowUsers(bson.M{"type": "user"}); err != nil {
		c.JSON(200, gin.H{"state": "no user match the condition"})
	} else {
		c.JSON(200, gin.H{"user": res})
	}
}

//localhost:8080/admin/delete
//BUG:delete a user dont exist return success...
func DelHandler(c *gin.Context) {
	log.Println(">>>Admin Delete User<<<")
	delId, _ := strconv.Atoi(c.PostForm("id"))
	filter := bson.M{"id": delId}
	if err := database.DeleteUser(filter); err != nil {
		c.JSON(200, gin.H{"state": "delete failed"})
		return
	}
	c.JSON(200, gin.H{"state": "delete success"})
}
