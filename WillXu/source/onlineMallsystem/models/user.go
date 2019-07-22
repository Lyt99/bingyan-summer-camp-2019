package models

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"onlineMallsystem/database"
	"onlineMallsystem/models/msg"
)

//new user sign up
func InsertUser(user msg.User) error {
	if _, err := database.UserColl.InsertOne(database.Ctx, bson.M{
		"username":            user.Username,
		"psw":                 user.Psw,
		"nickname":            user.Nickname,
		"mobile":              user.Mobile,
		"email":               user.Email,
		"total_view_count":    0,
		"total_collect_count": 0}); err != nil {
		return err
	}
	return nil
}

//find one user match the given filter
func FindUser(filter bson.M) (msg.User, error) {
	Msg := msg.User{}
	result := database.UserColl.FindOne(database.Ctx, filter)
	if err := result.Decode(&Msg); err != nil {
		return Msg, err
	}
	return Msg, nil
}

//update
func UserUpdate(id primitive.ObjectID, item string, i uint16) error {
	if _, err := database.UserColl.UpdateOne(database.Ctx,
		bson.M{"_id": id},
		bson.M{"$set": bson.M{item: i}}); err != nil {
		return err
	}
	return nil
}

//update user massage
func UpdateMsg(id primitive.ObjectID, item string, data string) error {
	if _, err := database.UserColl.UpdateOne(
		database.Ctx, bson.M{"_id": id},
		bson.M{"$set": bson.M{item: data}}); err != nil {
		return err
	}
	return nil
}
