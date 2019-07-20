package service

import (
	"final_task/model"
	"final_task/serializer"
	"gopkg.in/mgo.v2/bson"
	"regexp"
)

type UserRegisterService struct {
	UserName string `json:"user_name" bson:"user_name" form:"user_name"`
	Password string `json:"password"  bson:"password" form:"password"`
	Nickname string `json:"nickname"  bson:"nickname" form:"nickname"`
	Mobile   string `json:"mobile"    bson:"mobile"   form:"mobile"`
	Email    string `json:"email"     bson:"email"    form:"email"`
}

// 表单验证

func (service *UserRegisterService) Valid() *serializer.Response {
	s := model.GetMongoGlobalSession().Clone()
	defer s.Close()

	c := s.DB(model.DBName).C(model.CommodityCollectionName)

	count, err := c.Find(bson.M{
		"user_name": service.UserName,
	}).Count()
	if err != nil {
		return &serializer.Response{
			Success: false,
			Error:   "注册验证查询出错",
			Data:    nil,
		}
	}

	if count != 0 {
		return &serializer.Response{
			Success: false,
			Error:   "用户名已存在",
			Data:    nil,
		}
	}

	reg := regexp.MustCompile("^[1]([3-9])[0-9]{9}$")
	if reg.MatchString(service.Mobile) {
		return &serializer.Response{
			Success: false,
			Error:   "电话号码格式错误",
			Data:    nil,
		}
	}
	reg = regexp.MustCompile(`^[A-Za-z0-9\u4e00-\u9fa5]+@[a-zA-Z0-9_-]+(\.[a-zA-Z0-9_-]+)+$`)
	if reg.MatchString(service.Email) {
		return &serializer.Response{
			Success: false,
			Error:   "邮箱格式错误",
			Data:    nil,
		}
	}
	return nil
}

// 用户注册
func (service *UserRegisterService) Register() (model.User, *serializer.Response) {
	user := model.User{
		UserName: service.UserName,
		Nickname: service.Nickname,
		Mobile:   service.Mobile,
		Email:    service.Email,
	}

	if err := service.Valid(); err != nil {
		return user, err
	}

	if err := user.SetPassword(service.Password); err != nil {
		return user, &serializer.Response{
			Success: false,
			Error:   "密码加密失败",
			Data:    nil,
		}
	}

	// 创建用户
	s := model.GetMongoGlobalSession()
	s.Clone()
	defer s.Close()

	c := s.DB(model.DBName).C(model.CommodityCollectionName)
	if err := c.Insert(user); err != nil {
		return user, &serializer.Response{
			Success: false,
			Error:   "创建账户失败",
			Data:    nil,
		}
	}

	return user, nil
}
