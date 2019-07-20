package model

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"onlineMallsystem/conf/msg"
)

//new
func InsertCommodity(commodity msg.Commodity) error {
	if _, err := CommodityColl.InsertOne(ctx, bson.M{
		"pub_id":        commodity.PubUser,
		"title":         commodity.Title,
		"desc":          commodity.Desc,
		"category":      commodity.Category,
		"price":         commodity.Price,
		"picture":       commodity.Picture,
		"view_count":    0,
		"collect_count": 0}); err != nil {
		return err
	}
	return nil
}

//find one
func FindOneCommodity(filter bson.M) (msg.Commodity, error) {
	Msg := msg.Commodity{}
	result := CommodityColl.FindOne(ctx, filter)
	if err := result.Decode(&Msg); err != nil {
		return Msg, err
	}
	return Msg, nil
}

//find all
func FindAllCommodity(filter bson.M) ([]msg.MyCommodity, error) {
	cursor, err := CommodityColl.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	// iterate through all documents
	var res []msg.MyCommodity
	for cursor.Next(ctx) {
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
func GetCommoditiesList(page uint16, limit uint16, filter bson.M) ([]msg.ListCommodity, error) {
	cursor, err := CommodityColl.Find(ctx, filter, options.Find().SetSort(bson.M{"_id": -1}))
	if err != nil {
		return nil, err
	}
	// iterate through all documents
	var res []msg.ListCommodity
	min := page * limit
	max := min + limit
	for i := uint16(0); cursor.Next(ctx); i++ {
		for i >= min && i <= max {
			var p msg.ListCommodity
			// decode the document
			if err := cursor.Decode(&p); err != nil {
				return nil, err
			}
			res = append(res, p)
		}
		for i > max {
			return nil, err
		}
	}
	return res, nil
}

//update
func CommodityUpdate(id primitive.ObjectID, item string, i uint16) error {
	if _, err := CommodityColl.UpdateOne(ctx,
		bson.M{"_id": id},
		bson.M{"$set": bson.M{item: i}}); err != nil {
		return err
	}
	return nil
}

//delete one
func DeleteCommodity(filter bson.M) error {
	if _, err := CommodityColl.DeleteOne(ctx, filter); err != nil {
		return err
	}
	return nil
}
