package msg

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var IdentityKey = "id"

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

type Commodity struct {
	Id      primitive.ObjectID `bson:"_id"`
	PubUser string `bson:"pub_id"`

	Title    string  `json:"title" binding:"required"`
	Desc     string  `json:"desc" binding:"required"`
	Category uint16  `json:"category" binding:"required"`
	Price    float64 `json:"price"  binding:"required"`
	Picture  string  `json:"picture"  binding:"required"`

	ViewCount    uint16 `bson:"view_count"`
	CollectCount uint16 `bson:"collect_count"`
}

type SimpleCommodity struct {
	Id      primitive.ObjectID `bson:"_id"`
	Title    string  `json:"title" binding:"required"`
}