package controller

import (
	"MemberManager_TestPrograme/model"
	"MemberManager_TestPrograme/util/context"
	"crypto/md5"
	"encoding/hex"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"time"
)

type (
	Form struct {
		UserID      string `json:"User_ID"  form:"User_ID" query:"User_ID" validate:"required"`
		Name        string `json:"name"  form:"name" query:"name" validate:"required"`
		Password    string `json:"password"  form:"password" query:"password" validate:"required"`
		PhoneNumber string `json:"phone_number"  form:"phone_number" query:"phone_number" validate:"required,len=11"`
		Email       string `json:"email"  form:"email" query:"email" validate:"required,email"`
		IsAdmin     string `json:"is_admin"  form:"is_admin" query:"is_admin" validate:"required"`
	}
	jwtClaims struct {
		UserID string `json:"User_ID"`
		Admin  bool   `json:"admin"`
		jwt.StandardClaims
	}
)

func (f *Form) SendUserID() string {
	return f.UserID
}

func (f *Form) SendName() string {
	return f.Name
}

func (f *Form) SendPassword() string {
	// 使用md5进行加密后在传送
	h := md5.New()

	h.Write([]byte(f.Password))
	f.Password = hex.EncodeToString(h.Sum(nil))

	return f.Password
}

func (f *Form) SendPhoneNumber() string {
	return f.PhoneNumber
}

func (f *Form) SendEmail() string {
	return f.Email
}

func (f *Form) SendIsAdmin() bool {
	if f.IsAdmin == "I'm admin" {
		return true
	}

	return false
}

func CreateUser(c echo.Context) (err error) {
	form := new(Form)
	if err = c.Bind(form); err != nil {
		return context.Error(c, http.StatusBadRequest, "bad request", err)
	}
	err = c.Validate(*form)
	if err != nil {
		return context.Error(c, http.StatusBadRequest, "输入出错", err)

	}

	//

	user := model.User{}
	user.Name = form.SendName()
	user.UserID = form.SendUserID()
	user.PhoneNumber = form.SendPhoneNumber()
	user.Password = form.SendPassword()
	user.Email = form.SendEmail()
	user.IsAdmin = form.SendIsAdmin()
	if err = model.AddNewUser(user); err != nil {
		return context.Error(c, http.StatusBadRequest, "注册失败", err)
	}

	return context.Success(c, "完美")
}

func Login(c echo.Context) (err error) {
	UserID := c.FormValue("User_ID")
	password := c.FormValue("password")
	isExist, isAdmin, err := model.UserExists(UserID, password)
	if !isExist {
		return context.Error(c, http.StatusUnauthorized, "请输入正确的用户ID和密码", err)
	}

	// JWT
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = UserID
	if isAdmin {
		claims["admin"] = true
	} else {
		claims["admin"] = false
	}
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, bson.M{
		"token": t,
	})

}

func GetUserInfo(c echo.Context) (err error) {

	println("in GETUserInfo")
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	isAdmin := claims["admin"].(bool)
	userID := claims["id"].(string)
	queryID := c.FormValue("User_ID")
	println(userID, isAdmin)
	if userID != queryID && !isAdmin {
		return context.Error(c, http.StatusForbidden, "没有权限", err)
	}
	// var u model.User
	u, err := model.GetUserInfo(queryID)
	if err != nil {
		return context.Error(c, http.StatusNotFound, "没有找到哦", err)
	}

	return context.Success(c, bson.M{
		"UserID":      queryID,
		"PhoneNumber": u.PhoneNumber,
		"Email":       u.Email,
	})
}

func GetAllInfo(c echo.Context) (err error) {
	user := c.Get("admin").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	isAdmin := claims["admin"].(bool)

	if !isAdmin {
		return context.Error(c, http.StatusForbidden, "没有权限", err)
	}
	u, err := model.GetAllInfo()
	if err != nil {
		return context.Error(c, http.StatusNotFound, "查询炸了", err)
	}
	return c.JSON(http.StatusOK, u)
}

func DelUser(c echo.Context) (err error) {
	user := c.Get("admin").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	isAdmin := claims["admin"].(bool)
	queryID := c.FormValue("User_ID")

	if !isAdmin {
		return context.Error(c, http.StatusForbidden, "没有权限", err)
	}

	err = model.DeleteUser(queryID)
	if err != nil {
		return context.Error(c, http.StatusNotFound, "没有找到哦", err)
	}
	return context.Success(c, "完美")
}

func UpdateUserInfo(c echo.Context) (err error) {
	f := new(Form)
	if err = c.Bind(f); err != nil {
		return context.Error(c, http.StatusBadRequest, "bad request", err)
	}
	user := c.Get("admin").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	isAdmin := claims["admin"].(bool)
	userID := claims["id"].(string)
	queryID := f.UserID

	if queryID != userID && !isAdmin {
		return context.Error(c, http.StatusForbidden, "没有权限", nil)
	}

	err = c.Validate(*f)
	if err != nil {
		return context.Error(c, http.StatusBadRequest, "输入出错", err)

	}

	u := model.User{}
	u.Name = f.SendName()
	u.PhoneNumber = f.SendPhoneNumber()
	u.Password = f.SendPassword()
	u.IsAdmin = f.SendIsAdmin()
	if err = model.UpdateUser(f.UserID, u); err != nil {
		return context.Error(c, http.StatusBadRequest, "输入有误", err)
	}

	return context.Success(c, "完美")
}
