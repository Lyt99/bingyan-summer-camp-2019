package model

var IdentityKey = "id"
//user sign up message
type SignForm struct {
	Id uint16		`json:"id" form:"id" binding:"required"`
	Psw string		`json:"psw" form:"psw" binding:"required"`
	Name string		`json:"name" form:"name" binding:"required"`
	Tel string		`json:"tel" form:"tel" binding:"required"`
	Email string	`json:"email" form:"email" binding:"required"`
}

type LoginForm struct {
	Id uint16		`json:"id" form:"id" binding:"required"`
	Psw string		`json:"psw" form:"psw" binding:"required"`
}

//user update message
type UpdateForm struct {
	Item string		`json:"item" form:"item" binding:"required"`
	Context string	`json:"context" form:"context" binding:"required"`
}