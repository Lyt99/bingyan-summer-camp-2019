package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"onlineMallsystem/conf/Err"
	"onlineMallsystem/conf/msg"
	"onlineMallsystem/model"
)

//
func NewCommodity(c *gin.Context)  {
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
	strid:=fmt.Sprintf("%v",id)
	newCommodity.PubUser , _ =primitive.ObjectIDFromHex(strid)

	//insert commodity
	if err := model.InsertCommodity(newCommodity); err != nil {
		c.JSON(200, Err.InsertFailedJson)
		return
	}
	c.JSON(200, gin.H{
		"success": true,
		"error":"",
		"data":"ok"})
}

func DetailCommodity(c *gin.Context)  {
	
}