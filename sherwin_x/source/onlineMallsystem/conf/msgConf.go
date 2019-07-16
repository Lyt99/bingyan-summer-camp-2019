package conf

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	Id          primitive.ObjectID `bson:"_id"`
	CreatedTime time.Time          `bson:"created_at"`
	Visitor     uint16             `bson:"visitor"`

	UserType string `json:"type" form:"type" binding:"required"`
	Psw      string `json:"psw" form:"psw" binding:"required"`
	Name     string `json:"name" form:"name" binding:"required"`
	Tel      string `json:"tel" form:"tel" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required"`
}

type LoginForm struct {
	Id           primitive.ObjectID `bson:"_id"`
	UserType string `json:"type" form:"type" binding:"required"`
	Tel      string `json:"tel" form:"tel" binding:"required"`
	Psw      string `json:"psw" form:"psw" binding:"required"`
}

type Ware struct {
	Id           primitive.ObjectID `bson:"_id"`
	Name         string             `json:"name" form:"name" binding:"required"`
	Introduction string             `json:"introduction" form:"introduction" binding:"required"`
	Price        string             `json:"price" form:"price" binding:"required"`
	Quality      uint16             `json:"quality" form:"quality" binding:"required"`
	Tags         []string           `json:"tags" form:"tags" binding:"required"`
}
