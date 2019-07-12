package database

//source ~/.bashrc
//mongod --dbpath /home/sherwin/tools/mongodb/data

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func GetDatabase() *mongo.Client {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	//db = client.Database("demo")
	//userColl= client.Database("demo").Collection("user")
	return client
}
