package model

var IdentityKey = "id"

//user sign up message
/*{
	"id":1,
	"psw":"123",
	"name":"sherwin",
	"tel":"10086",
	"email":"18@16.com"
}*/
type SignForm struct {
	//DBid          primitive.ObjectID `bson:"_id"`
	Id    uint16 `json:"id" form:"id" binding:"required"`
	Psw   string `json:"psw" form:"psw" binding:"required"`
	Name  string `json:"name" form:"name" binding:"required"`
	Tel   string `json:"tel" form:"tel" binding:"required"`
	Email string `json:"email" form:"email" binding:"required"`
}

//user login message
type LoginForm struct {
	Id  uint16 `json:"id" form:"id" binding:"required"`
	Psw string `json:"psw" form:"psw" binding:"required"`
}

//user update message
type UpdateForm struct {
	Item    string `json:"item" form:"item" binding:"required"`
	Context string `json:"context" form:"context" binding:"required"`
}
