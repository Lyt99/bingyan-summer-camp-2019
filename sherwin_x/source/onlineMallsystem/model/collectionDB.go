package model

import (
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"onlineMallsystem/conf/msg"
)

//new collection
func InsertCollection(new msg.Collection) error {
	if _, err := CollectionColl.InsertOne(ctx, bson.M{
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
	result := CollectionColl.FindOne(ctx, filter)
	if err := result.Decode(&Msg); err != nil {
		return Msg, err
	}
	return Msg, nil
}

//find all
func FindAllCollection(filter bson.M) ([]msg.MyCollection, error) {
	cursor, err := CollectionColl.Find(ctx, filter)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	// iterate through all documents
	var res []msg.MyCollection
	for cursor.Next(ctx) {
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
	if _, err := CollectionColl.DeleteOne(ctx, filter); err != nil {
		return err
	}
	return nil
}
