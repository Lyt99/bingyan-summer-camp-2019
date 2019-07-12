package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type msg struct {
	ID string
	psw string
	name string
	tel string
	email string
}
var db *mongo.Database
var userColl *mongo.Collection
var ctx context.Context
func init(){
	ctx,_:=context.WithTimeout(context.Background(),10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
		fmt.Println("init failed")
	}
	db = client.Database("demo")
	userColl= client.Database("demo").Collection("user")
}

func main()  {
	r:=gin.Default()
	r.POST("/sign",SignHander)
	r.POST("/login",LogHandler)
	r.POST("/update",UpdateHandler)
	r.POST("/admin",AdminHandler)

	r.Run(":8080")
}

func SignHander(c *gin.Context)  {
	user := msg{}
	user.ID=c.PostForm("ID")
	user.psw=c.PostForm("password")
	user.name=c.PostForm("name")
	user.tel=c.PostForm("TEL")
	user.email=c.PostForm("e-mail")
	fmt.Println("submiting...")
	AddUser(user)
	c.JSON(200,gin.H{"html":"success"})
	return
	}
func AddUser(newUser msg) {
	if result, err := userColl.InsertOne(ctx, bson.M{
		"ID":    newUser.ID,
		"psw":   newUser.psw,
		"name":  newUser.name,
		"tel":   newUser.tel,
		"email": newUser.email}); err == nil {
		log.Println(result)
		fmt.Println("submit success！")

	} else {
		fmt.Print("submit failed")
	}
}

func LogHandler(c *gin.Context)  {
	user:=msg{}
	user.ID=c.PostForm("ID")
	user.psw=c.PostForm("password")
	
	result := userColl.FindOne(ctx, bson.M{
		"ID": user.ID,
		"psw":user.psw})
	if err := result.Decode(&user); err != nil {
		fmt.Println("login failed.")
	}else{
		fmt.Println("login success！")
		c.JSON(200,gin.H{"html":"success"})
	}
}

func UpdateHandler(c *gin.Context)  {
	data:=msg{}
	data.ID=c.PostForm("ID")
	data.psw=c.PostForm("password")
	data.name=c.PostForm("name")
	data.tel=c.PostForm("TEL")
	data.email=c.PostForm("e-mail")
	newdata:=msg{}
	newdata.ID=c.PostForm("new_ID")
	newdata.psw=c.PostForm("new_password")
	newdata.name=c.PostForm("new_name")
	newdata.tel=c.PostForm("new_TEL")
	newdata.email=c.PostForm("new_e-mail")
	if result, err := userColl.UpdateOne(
		ctx, bson.M{"ID":data.ID},
		bson.M{"$set": bson.M{"ID": newdata.ID}}); err == nil {
		log.Println(result)
		fmt.Println(data)
		fmt.Println(newdata)
	} else {
		log.Fatal(err)
	}
	if result, err := userColl.UpdateOne(
		ctx, bson.M{"psw":data.psw},
		bson.M{"$set": bson.M{"psw": newdata.psw}}); err == nil {
		log.Println(result)
		fmt.Println(data)
		fmt.Println(newdata)
	} else {
		log.Fatal(err)
	}
	if result, err := userColl.UpdateOne(
		ctx, bson.M{"name":data.name},
		bson.M{"$set": bson.M{"name": newdata.name}}); err == nil {
		log.Println(result)
		fmt.Println(data)
		fmt.Println(newdata)
	} else {
		log.Fatal(err)
	}
	if result, err := userColl.UpdateOne(
		ctx, bson.M{"tel":data.tel},
		bson.M{"$set": bson.M{"tel": newdata.tel}}); err == nil {
		log.Println(result)
		fmt.Println(data)
		fmt.Println(newdata)
	} else {
		log.Fatal(err)
	}
	if result, err := userColl.UpdateOne(
		ctx, bson.M{"email":data.email},
		bson.M{"$set": bson.M{"email": newdata.email}}); err == nil {
		log.Println(result)
		fmt.Println(data)
		fmt.Println(newdata)
	} else {
		log.Fatal(err)
	}
}

func AdminHandler(c *gin.Context)  {

}