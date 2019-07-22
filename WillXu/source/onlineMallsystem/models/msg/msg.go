package msg

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var IdentityKey = "id"

//user message
type User struct {
	Id primitive.ObjectID `bson:"_id"`

	Username string `json:"username" form:"username" binding:"required"`
	Psw      string `json:"password" form:"password" binding:"required"`
	Nickname string `json:"nickname" form:"nickname" binding:"required"`
	Mobile   string `json:"mobile" form:"mobile" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required"`

	TotalViewCount    uint16 `bson:"total_view_count"`
	TotalCollectCount uint16 `bson:"total_collect_count"`
}

type LoginForm struct {
	Id primitive.ObjectID `bson:"_id"`

	Username string `json:"username" form:"username" binding:"required"`
	Psw      string `json:"password" form:"password" binding:"required"`
}

//commodity
type Commodity struct {
	Id      primitive.ObjectID `bson:"_id"`
	PubUser string             `bson:"pub_id"`

	Title    string  `json:"title" binding:"required"`
	Desc     string  `json:"desc" binding:"required"`
	Category uint16  `json:"category" binding:"required"`
	Price    float64 `json:"price" binding:"required"`
	Picture  string  `json:"picture" binding:"required"`

	ViewCount    uint16 `bson:"view_count"`
	CollectCount uint16 `bson:"collect_count"`
}

type MyCommodity struct {
	Id    primitive.ObjectID `bson:"_id"`
	Title string             ` bson:"title" `
}

type GetCommodity struct {
	Page     int    `json:"page" binding:"required"`
	Limit    int    `json:"limit" binding:"required"`
	Category int    `json:"category" binding:"required"`
	Keyword  string `json:"keyword" binding:"required"`
}

type ListCommodity struct {
	Id       primitive.ObjectID `bson:"_id"`
	Title    string             `bson:"title"`
	Desc     string             `bson:"desc"`
	Category uint16             `bson:"category"`
	Price    float64            `bson:"price"`
	Picture  string             `bson:"picture"`
}

//collection
type Collection struct {
	UserId string `json:"user_id" bson:"user_id" `
	Id     string `json:"id" bson:"id" binding:"required"`
	Title  string `json:"title" bson:"title" `
}

type MyCollection struct {
	Id    string ` bson:"id"`
	Title string ` bson:"title" `
}

//keyword
type Key struct {
	Keyword string `bson:"keyword"`
	Count   uint16 `bson:"count"`
}

type GetKey struct {
	Keyword string `bson:"keyword"`
}
