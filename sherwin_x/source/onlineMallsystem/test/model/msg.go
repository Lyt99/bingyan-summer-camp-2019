package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

//user sign up message
type UserMsg struct {
	Id         primitive.ObjectID `json:"id" form:"id"`
	Createtime time.Time          `json:"createtime"`
	Updatetime time.Time          `json:"updatetime"`
	Visitor    uint16             `json:"visitor"`

	Tel      string `json:"tel" form:"tel" binding:"required"`
	Psw      string `json:"psw" form:"psw" binding:"required"`
	Name     string `json:"psw" form:"psw" binding:"required"`
	Usertype string `json:"usertype" form:"usertype" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required"`
}
