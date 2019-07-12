package controller

import (
	"awesomeProject/model"
	"fmt"
	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"strconv"
)
//admin authority
func AdminCallback(c *gin.Context)(interface{},error){
	log.Println(">>>Admin Authoring<<<")
	user:=model.LoginForm{}
	_ =c.ShouldBind(&user)
	if user.Id==2019&&user.Psw=="0711"{
	return "200", nil
	}else {
	return nil, jwt.ErrFailedAuthentication}
}

// auth test
func HelloAdminHandler(c *gin.Context)  {
	log.Println(">>>Admin Auth Test<<<")
	c.JSON(200,gin.H{
		"ID":"admin",
		"text":"Do What You Want!",
	})
}

//localhost:8080/admin/find
func FindHandler(c *gin.Context)  {
	log.Println(">>>Admin Find User<<<")
	var p model.SignForm
	findId, _ :=strconv.Atoi(c.PostForm("id"))
	result := UserColl.FindOne(ctx, bson.M{"id": findId})
	if err := result.Decode(&p); err != nil {
		log.Println("not found")
	}else {
		fmt.Printf("post: %+v\n", p)
	}
}

//localhost:8080/admin/show
func ShowHandler(c *gin.Context){
	log.Println(">>>Admin Show All Users<<<")
	//filter := bson.M{"tags": bson.M{"$elemMatch": bson.M{"type": "user"}}}
	filter :=bson.M{"type": "user"}

	// find all documents
	cursor, err := UserColl.Find(ctx, filter)
	if err != nil {
		log.Println("err1")
		//log.Fatal(err)
	}

	// iterate through all documents
	for cursor.Next(ctx) {
		var p model.SignForm
		// decode the document
		if err := cursor.Decode(&p); err != nil {
			log.Println("err2")
			//log.Fatal(err)
		}
		fmt.Printf("UserMassage:%+v\n", p)
	}
}

//localhost:8080/admin/delete
func DelHandler(c *gin.Context)  {
	delId, _ :=strconv.Atoi(c.PostForm("id"))//change PostForm into int
	log.Println(">>>Admin Delete User<<<")
	if result, err := UserColl.DeleteOne(ctx, bson.M{"id": delId}); err != nil {
		log.Println("delete failed")
	} else {
		log.Println(result)
		log.Println("deleted user:",delId)
	}
}