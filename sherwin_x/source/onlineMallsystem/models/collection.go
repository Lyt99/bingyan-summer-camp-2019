package models

import (
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"onlineMallsystem/database"
	"onlineMallsystem/models/msg"
)

//new collection
func InsertCollection(new msg.Collection) error {
	if _, err := database.CollectionColl.InsertOne(database.Ctx, bson.M{
		"user_id": new.UserId,
		"id":      new.Id,
		"title":   new.Title,
	}); err != nil {
		return err
	}
	return nil
}

//find one
func FindOneCollection(filter bson.M) (msg.Collection, error) {
	Msg := msg.Collection{}
	result := database.CollectionColl.FindOne(database.Ctx, filter)
	if err := result.Decode(&Msg); err != nil {
		return Msg, err
	}
	return Msg, nil
}

//find all
func FindAllCollection(filter bson.M) ([]msg.MyCollection, error) {
	cursor, err := database.CollectionColl.Find(database.Ctx, filter)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	// iterate through all documents
	var res []msg.MyCollection
	for cursor.Next(database.Ctx) {
		var p msg.MyCollection
		// decode the document
		if err := cursor.Decode(&p); err != nil {
			log.Println(err)
			return nil, err
		}
		log.Println(p)
		res = append(res, p)
	}
	return res, nil
}

//delete one
func DeleteOneCollection(filter bson.M) error {
	if _, err := database.CollectionColl.DeleteOne(database.Ctx, filter); err != nil {
		return err
	}
	return nil
}
