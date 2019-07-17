package model

type (
	User struct {
		UserID      string `json:"User_ID" bson:"User_ID" form:"User_ID" query:"User_ID"`
		Name        string `json:"name" bson:"name" form:"name" query:"User_ID"`
		Password    string `json:"password" bson:"password" form:"name" query:"name"`
		PhoneNumber string `json:"phone_number" bson:"phone_number" form:"name" query:"name"`
		Email       string `json:"email" bson:"email" form:"email" query:"email"`
		IsAdmin     bool   `json:"is_admin" bson:"is_admin" form:"is_admin" query:"is_admin"`
	}
)

type UserInfoSender interface {
	SendUserID() string
	SendName() string
	SendPassword() string
	SendPhoneNumber() string
	SendEmail() string
	SendIsAdmin() bool
}
