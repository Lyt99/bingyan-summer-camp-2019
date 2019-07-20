package model

import (
	"gopkg.in/mgo.v2"
	"log"
)

var mongoSession *mgo.Session

const (
	DBName                  = "FinalTask"
	UserCollectionName      = "users"
	CommodityCollectionName = "commodities"
	url                     = "localhost:27017"
)

func init() {
	log.Println("Init database......")

	session, err := getMongoSession()
	if err != nil {
		log.Println("MongoDB init error")
		log.Panic(err)
		return
	}
	mongoSession = session

	log.Println("Database init done")
}

func getMongoSession() (*mgo.Session, error) {
	mgoSession, err := mgo.Dial(url)
	if err != nil {
		log.Println("MongoDb Dial error")
		log.Panic(err)
		return nil, err
	}

	mgoSession.SetMode(mgo.Monotonic, true)

	return mgoSession, nil
}

func GetMongoGlobalSession() *mgo.Session {
	return mongoSession
}
