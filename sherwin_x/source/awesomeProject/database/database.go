package database

import (
	"awesomeProject/model"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var UserColl *mongo.Collection
var ctx context.Context

//connect database
func init() {
	log.Println(">>>Database Connecting<<<")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	UserColl=client.Database("demo").Collection("user")
}

//new user sign in
func InsertUser(user model.SignForm)error  {
	if _, err := UserColl.InsertOne(ctx, bson.M{
		"type":  "user",
		"id":    user.Id,
		"psw":   user.Psw,
		"name":  user.Name,
		"tel":   user.Tel,
		"email": user.Email}); err != nil {
		return err
	}
	return nil
}

//find user massage with filter
func FindUser(filter bson.M) (model.SignForm,error) {
	Msg:=model.SignForm{}
	result := UserColl.FindOne(ctx, filter)
	if err := result.Decode(&Msg); err != nil {
		return Msg,err
	}
	return Msg,nil
}

//update user massage
func UpdateMsg(id interface{},newdate model.UpdateForm) error {
	 result, _ := UserColl.UpdateOne(
		ctx, bson.M{"id": id},
		bson.M{"$set": bson.M{newdate.Item: newdate.Context}})
	if result.ModifiedCount == 0 {
		return errors.New("not change")
	}
	return nil
}
