package service

import "MemberManager_TestPrograme/model"

type UserLoginService struct {
	UserName string `form:"user_ame" json:"user_name" bson:"user_name"`
	Password string `form:"password" json:"password" bson:"password"`
}

func (u *UserLoginService) Login() (model.User, *serializer.Resonse) {

}
