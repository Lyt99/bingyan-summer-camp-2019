package test

import (
	"awesomeProject/model"
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

var UserColl *mongo.Collection
var ctx context.Context

func SignHandler(c *gin.Context) {
	log.Println(">>>Message Submitting<<<")
	newuser := model.SignForm{}

	//bind sign massage
	if err := c.ShouldBind(&newuser); err != nil {
		c.JSON(400, gin.H{"warning": "invalid massage"})
		return
	}

	//check id availability
	if err := idCheck(newuser); err == nil {
		c.JSON(400, gin.H{"warning": "invalid id!"})
	}

}

func idCheck(userMsg model.SignForm) error {
	log.Println(">>>Id Checking<<<")
	newUser := userMsg
	result := UserColl.FindOne(ctx, bson.M{"id": newUser.Id})
	if err := result.Decode(&newUser); err != nil {
		return err
	} else {
		return nil
	}
}