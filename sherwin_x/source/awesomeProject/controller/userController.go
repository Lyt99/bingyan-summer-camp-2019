package controller

import (
	"awesomeProject/database"
	"awesomeProject/model"
	"context"
	"crypto/md5"
	"fmt"
	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

var UserColl *mongo.Collection
var ctx context.Context

//connect DataBase
func init()  {
	log.Println(">>>Database Connecting<<<")
	UserColl=database.GetDatabase().Database("demo").Collection("user")
}

//localhost:8080/sign　　
func SignHandler(c *gin.Context)  {
	log.Println(">>>Message Submitting<<<")
	newuser:=model.SignForm{}
	if err:=c.ShouldBind(&newuser);err!=nil{
		c.JSON(400,gin.H{"statu":"invalid massage"})
	}else{
		if err:=idCheck(newuser);err==nil{
			c.JSON(400,gin.H{"statu":"invalid id!"})
		}else {
			newuser.Psw=encode(newuser.Psw)//encode psw to md5
			if err:=AddUser(newuser);err!=nil{
				c.JSON(500,gin.H{"statu":"sorry, sign up failed!"})
			}else{c.JSON(200,gin.H{"statu":"sign up success!"})}
		}
	}
}

func idCheck(userMsg model.SignForm) error {
	log.Println(">>>Id Checking<<<")
	newUser := userMsg
	result := UserColl.FindOne(ctx, bson.M{"id": newUser.Id})
	if err := result.Decode(&newUser); err != nil {
		return err
	}else {return nil}
}
func encode(psw string) string {
	pswMd5:=md5.New()
	pswMd5.Write([]byte(psw))
	psw=string(pswMd5.Sum(nil))
	return psw
}
func AddUser(newUser model.SignForm) error {
	log.Println(">>>Message Storing<<<")
	if result, err := UserColl.InsertOne(ctx, bson.M{
		"type":"user",
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
func UserCallback(c *gin.Context) (interface{},error) {
	log.Println(">>>User Authoring<<<")
	user:=model.LoginForm{}
	_ = c.ShouldBind(&user)
	user.Psw=encode(user.Psw)
	result := UserColl.FindOne(ctx, bson.M{
		"id":user.Id,
		"psw":user.Psw,})
	if err := result.Decode(&user); err != nil {
		log.Println("user login failed")
		return nil, jwt.ErrFailedAuthentication
	} else {
		log.Println("user been login")
		//c.JSON(200,gin.H{"state":"success"})
		return &model.LoginForm{
			Id:user.Id,}, nil}
}

//localhost:8080/user/hello
func HelloUserHandler(c *gin.Context)  {
	log.Println(">>>User Auth Test<<<")
	id,err:=c.Get(model.IdentityKey)
	if !err{
		log.Println("id_get_failed")
		fmt.Println(id)
	}
	c.JSON(200,gin.H{
		"userID":id,
		"text":"Welcome!",
	})
}

//localhost:8080/user/update　　　
func UpdateHandler(c *gin.Context)  {
	log.Println(">>>User Message Update<<<")
	id,err:=c.Get(model.IdentityKey)
	if !err{
		log.Println("id_get_failed")
		fmt.Println(id)
	}
	//fmt.Println(id)
	newdate:=model.UpdateForm{}
	_ =c.ShouldBind(&newdate)
	if newdate.Item=="psw"||newdate.Item=="name"||newdate.Item=="tel"||newdate.Item=="email" {
		if result, err := UserColl.UpdateOne(
			ctx, bson.M{"id": id},
			bson.M{"$set": bson.M{newdate.Item: newdate.Context}}); err != nil { //id is from token,so err must != nil
			log.Fatal(err)
		} else {
			log.Println(result)
			if result.ModifiedCount == 1 {
				c.JSON(200, gin.H{"statu": "update success!"})
			} else {
				c.JSON(400, gin.H{"statu": "massage not change"})
			}
		}
	}else {
		c.JSON(400,gin.H{"statu":"invalid item"})
	}
}