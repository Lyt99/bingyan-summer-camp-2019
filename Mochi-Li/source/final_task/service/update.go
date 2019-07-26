package service

import (
	"final_task/model"
	"final_task/serializer"
	"gopkg.in/mgo.v2/bson"
)

type UserUpdateService struct {
	Password string `json:"password" bson:"password" form:"password"`
	Nickname string `json:"nickname" bson:"nickname" form:"nickname"`
	Mobile   string `json:"mobile"   bson:"mobile"   form:"mobile"`
	Email    string `json:"email"    bson:"email"   form:"email"`
}

func (service *UserUpdateService) Updater(username string) *serializer.Response {
	s := model.GetMongoGlobalSession().Copy()
	defer s.Close()

	u := model.User{
		UserName: username,
	}
	if err := u.FindUserAsName(); err != nil {
		return &serializer.Response{
			Success: false,
			Error:   "查询出错",
			Data:    nil,
		}
	}

	u.Nickname = service.Nickname
	u.Mobile = service.Mobile
	u.Email = service.Email
	if service.Password != "" {
		if err := u.SetPassword(service.Password); err != nil {
			return &serializer.Response{
				Success: false,
				Error:   "设置密码失败",
				Data:    nil,
			}
		}
	}
	if err := s.DB(model.DBName).C(model.UserCollectionName).Update(bson.M{"username": username}, u); err != nil {
		return &serializer.Response{
			Success: false,
			Error:   "更新用户失败",
			Data:    nil,
		}
	}
	return &serializer.Response{
		Success: true,
		Error:   "",
		Data:    "ok",
	}
}
