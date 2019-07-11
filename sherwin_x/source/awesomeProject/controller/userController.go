package controller

import (
	"awesomeProject/database"
	"awesomeProject/model"
	"context"
	"fmt"
	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)



var UserColl *mongo.Collection
var ctx context.Context

func init()  {
	log.Println(">>>Connecting Database<<<")
	UserColl=database.GetDatabase().Database("demo").Collection("user")
}


//sign up
func SignHandler(c *gin.Context)  {
	log.Println(">>>Submitting Message<<<")
	newuser:=model.SignForm{}
	_ =c.ShouldBind(&newuser)
	if err:=idCheck(newuser);err==nil{
		c.JSON(400,gin.H{"statu":"invalid id!"})
	}else {
		if err:=AddUser(newuser);err!=nil{
		c.JSON(500,gin.H{"statu":"sorry, sign up failed!"})
		}else{c.JSON(200,gin.H{"statu":"sign up success!"})}
	}
}

func idCheck(userMsg model.SignForm) error {
	newUser := userMsg
	result := UserColl.FindOne(ctx, bson.M{"id": newUser.Id})
	if err := result.Decode(&newUser); err != nil {
		return err
	}else {return nil}
}

func AddUser(newUser model.SignForm) error {
	log.Println(">>>Storing Message<<<")
	if result, err := UserColl.InsertOne(ctx, bson.M{
		"id":  newUser.Id,
		"psw": newUser.Psw,
		"name":  newUser.Name,
		"tel":   newUser.Tel,
		"email": newUser.Email}); err == nil {
			log.Println("you've got a new user")
			log.Println(result)
		return nil
	} else {return err}
}


//user authority
//To be refactored
func UserCallback(c *gin.Context) (interface{},error) {
	log.Println(">>>User Authoring<<<")
	user:=model.LoginForm{}
	_ = c.ShouldBindUri(&user)
	result := UserColl.FindOne(ctx, bson.M{
		"id":user.Id,
		"psw":user.Psw,})
	if err := result.Decode(&user); err != nil {
		fmt.Println("user login failed")
		return nil, jwt.ErrFailedAuthentication
	} else {
		fmt.Println("user been login")
		c.JSON(200,gin.H{"state":"success"})
		return &model.LoginForm{
			Id:user.Id,}, nil}
}

//auth test
func HelloUserHandler(c *gin.Context)  {
	log.Println(">>>User Auth Test<<<")
	c.JSON(200,gin.H{
		"ID":"user",
		"text":"Welcome!",
	})
}

// up data
//TBD
func UpdateHandler(c *gin.Context)  {
	log.Println(">>>User Message Updating<<<")
	id,err:=c.Get(model.IdentityKey)
	if !err{
		log.Println("id_get_failed")
		fmt.Println(id)
	}
	//fmt.Println(id)
	newdate:=model.UpdateForm{}
	_ =c.ShouldBindUri(&newdate)
	if result, err := UserColl.UpdateOne(
		ctx, bson.M{"ID":id},
		bson.M{"$set": bson.M{newdate.Item: newdate.Context}}); err == nil {
		log.Println(result)
		log.Println(id)
	} else {
		log.Fatal(err)
	}
}