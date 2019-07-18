package model

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var UserColl *mongo.Collection
var CommodityColl *mongo.Collection
var CollectionColl *mongo.Collection
var ctx context.Context

//connect database
func init() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	UserColl = client.Database("mall").Collection("user")
	CommodityColl = client.Database("mall").Collection("commodity")
	CollectionColl = client.Database("mall").Collection("collection")
}
