package model

import (
	"gopkg.in/mgo.v2"
	"log"
)

var (
	mongoSession *mgo.Session
)

const (
	DBName             = "MemberManager"
	userCollectionName = "users"
	url                = "localhost:27017"
)

func init() {
	log.Println("Init database...")

	session, err := getMongoSession()
	if err != nil {
		log.Println("MongoDB init error!")
		log.Panic(err)
		return
	}
	mongoSession = session

	log.Println("Database init done!")

	return
}

func getMongoSession() (*mgo.Session, error) {
	mgosession, err := mgo.Dial(url)
	if err != nil {
		log.Println("MongoDB dial error")
		log.Panic(err)
		return nil, err
	}

	mgosession.SetMode(mgo.Monotonic, true)

	return mgosession, nil
}

func GetMongoGlobalSession() *mgo.Session {
	return mongoSession
}
