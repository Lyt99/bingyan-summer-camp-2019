package main

import (
	"MemberManager_TestPrograme/controller"
	m "MemberManager_TestPrograme/middleware"
	_ "MemberManager_TestPrograme/model"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

const (
	SignUpURL           = "/sign-up"
	LoginURL            = "/login-in"
	UserGetUserInfoURL  = "/User/get"
	AdminGetUserInfoURL = "/User/get"
	DelUserURL          = "/User/delete"
	UpdateURL           = "/User/update"
	GetAllUerInfo       = "/User/get-all"
)

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Validator = m.GetValidator()
	// Route => handler





	e.POST(SignUpURL, controller.CreateUser)
	e.POST(LoginURL, controller.Login)

	r :=e.Group("/User")
	r.Use(middleware.JWT([]byte("secret")))
	e.POST(UserGetUserInfoURL, controller.GetUserInfo)
	e.POST(AdminGetUserInfoURL, controller.GetUserInfo)
	e.POST(DelUserURL, controller.DelUser)
	e.POST(UpdateURL, controller.UpdateUserInfo)
	e.POST(GetAllUerInfo, controller.GetAllInfo)

	e.Logger.Fatal(e.Start(":2333"))
}
