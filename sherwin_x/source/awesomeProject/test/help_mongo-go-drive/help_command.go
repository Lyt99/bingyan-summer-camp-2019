package main //mgo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type User struct {
	ID    primitive.ObjectID `bson:"_id"`
	Name  string             `bson:"dbname",json:"jsonname"`
	Phone string
}

func main() {

	log.SetFlags(log.Lshortfile)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	db := client.Database("demo")

	// Drop database
	err = db.Drop(ctx)
	if err != nil {
		log.Fatal(err)
	}

	userColl := db.Collection("user")

	// Insert one
	if result, err := userColl.InsertOne(ctx, User{Name: "UserName", Phone: "1234567890"}); err == nil {
		log.Println(result)
	} else {
		log.Fatal(err)
	}

	// Insert many
	{
		users := []interface{}{
			User{Name: "UserName_0", Phone: "123"},
			User{Name: "UserName_1", Phone: "456"},
			User{Name: "UserName_2", Phone: "789"},
		}
		if result, err := userColl.InsertMany(ctx, users); err == nil {
			log.Println(result)
		} else {
			log.Fatal(err)
		}
	}

	// Find one
	{
		result := userColl.FindOne(ctx, bson.M{"phone": "1234567890"})
		var user User
		if err := result.Decode(&user); err != nil {
			log.Fatal(err)
		}
		log.Println(user)
	}

	// Find
	if cur, err := userColl.Find(ctx, bson.M{"phone": primitive.Regex{Pattern: "456", Options: ""}}); err == nil {
		defer cur.Close(ctx)
		for cur.Next(ctx) {
			var user User
			if err := cur.Decode(&user); err != nil {
				log.Fatal(err)
			}
			log.Println(user)
		}
	} else {
		log.Fatal(err)
	}

	// Update one
	if result, err := userColl.UpdateOne(
		ctx, bson.M{"phone": "123"},
		bson.M{"$set": bson.M{"dbname": "UserName_changed"}}); err == nil {
		log.Println(result)
	} else {
		log.Fatal(err)
	}

	// Update many
	if result, err := userColl.UpdateMany(
		ctx, bson.M{"phone": primitive.Regex{Pattern: "456", Options: ""}},
		bson.M{"$set": bson.M{"dbname": "UserName_changed"}}); err == nil {
		log.Println(result)
	} else {
		log.Fatal(err)
	}

	// Replace one
	{
		user := User{Name: "UserName_2_replaced", Phone: "789"}
		if result, err := userColl.ReplaceOne(ctx, bson.M{"phone": "789"}, user); err == nil {
			log.Println(result)
		} else {
			log.Fatal(err)
		}
	}

	// Delete one
	if result, err := userColl.DeleteOne(ctx, bson.M{"phone": "123"}); err == nil {
		log.Println(result)
	} else {
		log.Fatal(err)
	}

	// Delete many
	if result, err := userColl.DeleteMany(ctx, bson.M{"phone": primitive.Regex{Pattern: "456", Options: ""}}); err == nil {
		log.Println(result)
	} else {
		log.Fatal(err)
	}
}
