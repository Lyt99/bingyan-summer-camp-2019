package model

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"onlineMallsystem/conf/msg"
)

//new user sign in
func InsertUser(user msg.User) error {
	if _, err := UserColl.InsertOne(ctx, bson.M{
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
	result := UserColl.FindOne(ctx, filter)
	if err := result.Decode(&Msg); err != nil {
		return Msg, err
	}
	return Msg, nil
}

//update
func UserUpdate(id primitive.ObjectID, item string, i uint16) error {
	if _, err := UserColl.UpdateOne(ctx,
		bson.M{"_id": id},
		bson.M{"$set": bson.M{item: i}}); err != nil {
		return err
	}
	return nil
}
