package model

import (
	"go.mongodb.org/mongo-driver/bson"
	"onlineMallsystem/conf/msg"
)

//new
func InsertCommodity(commodity msg.Commodity) error {
	if _, err := CommodityColl.InsertOne(ctx, bson.M{
		"pub_id":commodity.PubUser,
		"title":commodity.Title,
		"desc":commodity.Desc,
		"category":commodity.Category,
		"price":commodity.Price,
		"picture":commodity.Picture,
		"view_count":0,
		"collect_count":0}); err != nil {
		return err
	}
	return nil
}
