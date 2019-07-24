package models

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"onlineMallsystem/database"
	"onlineMallsystem/models/msg"
	"time"
)

//new
func InsertCommodity(commodity msg.Commodity) error {
	if _, err := database.CommodityColl.InsertOne(database.Ctx, bson.M{
		"pub_id":        commodity.PubUser,
		"title":         commodity.Title,
		"desc":          commodity.Desc,
		"category":      commodity.Category,
		"price":         commodity.Price,
		"picture":       commodity.Picture,
		"create_time":   time.Now().Unix(),
		"type":          "commodity",
		"view_count":    0,
		"collect_count": 0}); err != nil {
		return err
	}
	return nil
}

//find one
func FindOneCommodity(filter bson.M) (msg.Commodity, error) {
	Msg := msg.Commodity{}
	result := database.CommodityColl.FindOne(database.Ctx, filter)
	if err := result.Decode(&Msg); err != nil {
		return Msg, err
	}
	return Msg, nil
}

//find all
func FindAllCommodity(filter bson.M) ([]msg.MyCommodity, error) {
	cursor, err := database.CommodityColl.Find(database.Ctx, filter)
	if err != nil {
		return nil, err
	}
	// iterate through all documents
	var res []msg.MyCommodity
	for cursor.Next(database.Ctx) {
		var p msg.MyCommodity
		// decode the document
		if err := cursor.Decode(&p); err != nil {
			return nil, err
		}
		res = append(res, p)
	}
	return res, nil
}

//get commodities list
func GetCommoditiesList(page int, limit int, filter bson.M) ([]msg.ListCommodity, error) {
	cursor, err := database.CommodityColl.Find(database.Ctx, filter, options.Find().SetSort(bson.M{"create_time": -1}))
	if err != nil {
		return nil, err
	}
	// iterate through all documents
	min := page * limit
	max := min + limit
	var res []msg.ListCommodity
	for i := 0; cursor.Next(database.Ctx); i++ {
		var p msg.ListCommodity
		// decode the document
		if err := cursor.Decode(&p); err != nil {
			return nil, err
		}
		if i >= min && i < max {
			res = append(res, p)
		}
	}
	return res, nil
}

//update
func CommodityUpdate(id primitive.ObjectID, item string, i uint16) error {
	if _, err := database.CommodityColl.UpdateOne(database.Ctx,
		bson.M{"_id": id},
		bson.M{"$set": bson.M{item: i}}); err != nil {
		return err
	}
	return nil
}

//delete one
func DeleteCommodity(filter bson.M) error {
	if _, err := database.CommodityColl.DeleteOne(database.Ctx, filter); err != nil {
		return err
	}
	return nil
}
