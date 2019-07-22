package models

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"onlineMallsystem/database"
	"onlineMallsystem/models/msg"
)

func KeyFunc(key msg.Key) error {
	//find key
	res, err := FindOneKeyword(key.Keyword)
	//if not found, insert one
	if err != nil {
		log.Println("NotFound:", err)
		err = InsertKeyword(key)
		if err != nil {
			return err
		}
		return nil
	}
	//if found, count+1
	if err = UpdateKeyword(res.Keyword, res.Count); err != nil {
		return err
	}
	return nil
}

//insert one
func InsertKeyword(keyword msg.Key) error {
	if _, err := database.KeywordColl.InsertOne(database.Ctx, bson.M{
		"type":    "key",
		"keyword": keyword.Keyword,
		"count":   keyword.Count,
	}); err != nil {
		return err
	}
	return nil
}

//find one
func FindOneKeyword(key string) (msg.Key, error) {
	var res msg.Key
	result := database.KeywordColl.FindOne(database.Ctx, bson.M{"keyword": key})
	if err := result.Decode(&res); err != nil {
		return res, err
	}
	return res, nil
}

//find key:if find,then count+1,else insert one
func UpdateKeyword(key string, count uint16) error {
	if _, err := database.KeywordColl.UpdateOne(database.Ctx,
		bson.M{"keyword": key},
		bson.M{"$set": bson.M{"count": count + 1}}); err != nil {
		return err
	}
	return nil
}

//find all
func FindAllKeyword() ([]string, error) {
	cursor, err := database.KeywordColl.Find(database.Ctx, bson.M{"type": "key"}, options.Find().SetSort(bson.M{"count": -1}))
	if err != nil {
		return nil, err
	}
	// iterate through all documents
	var res []string
	for i := 0; cursor.Next(database.Ctx); i++ {
		var p msg.GetKey
		// decode the document
		if err := cursor.Decode(&p); err != nil {
			return nil, err
		}
		if i >= 0 && i < 10 {
			res = append(res, p.Keyword)
			log.Println(res)
		}
	}
	return res, nil
}
