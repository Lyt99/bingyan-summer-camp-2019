package model

import (
	"go.mongodb.org/mongo-driver/bson"
	"onlineMallsystem/conf"
	"time"
)

//new user sign in
func InsertUser(user conf.User) error {
	if _, err := UserColl.InsertOne(ctx, bson.M{
		"created_at": time.Now(),
		"visitor":    0,
		"type":       user.UserType,
		"psw":        user.Psw,
		"name":       user.Name,
		"tel":        user.Tel,
		"email":      user.Email,}); err != nil {
		return err
	}
	return nil
}

//find one user match the given filter
func FindUser(filter bson.M) (conf.User, error) {
	Msg := conf.User{}
	result := UserColl.FindOne(ctx, filter)
	if err := result.Decode(&Msg); err != nil {
		return Msg, err
	}
	return Msg, nil
}
