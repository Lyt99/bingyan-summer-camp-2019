package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var UserColl *mongo.Collection
var ctx context.Context

//connect DataBase
func init() {
	log.Println(">>>Database Connecting<<<")
	UserColl = GetDatabase().Database("demo").Collection("user")
}

func GetDatabase() *mongo.Client {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	return client
}
