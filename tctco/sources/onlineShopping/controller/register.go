package controller

import (
	"github.com/gin-gonic/gin"
	"onlineShopping/model"
	"regexp"
)

type Response map[string]interface{}

func Register(c *gin.Context) {
	var json model.RegisterJSON
	response := ResponseInit()
	if err := c.ShouldBindJSON(&json); err != nil || json.Username == "login" {
		response["error"] = "输入信息不合法"
		c.JSON(400, response)
		return
	}
	if fmtCheck(json, c) && checkUserNotExist(json.Username, c) {
		if err := model.DBRegister(json); err != nil {
			dbError(c)
			return
		}
		response["success"] = true
		response["error"] = ""
		response["data"] = "ok"
		c.JSON(200, response)
	}

}

func fmtCheck(json model.RegisterJSON, c *gin.Context) bool {
	response := make(Response)
	response["success"] = false
	response["data"] = nil
	if !fmtCheckPassword(json.Password) {
		response["error"] = "密码长度不足"
		c.JSON(400, response)
		return false
	}

	if !fmtCheckEmail(json.Email) {
		response["error"] = "邮箱格式错误"
		c.JSON(400, response)
		return false
	}

	if !fmtCheckUsername(json.Username) {
		response["error"] = "用户名过短"
		c.JSON(400, response)
		return false
	}

	if !fmtCheckTelephone(json.Mobile) {
		response["error"] = "手机号格式错误"
		c.JSON(400, response)
		return false
	}

	if !fmtCheckUsername(json.Nickname) {
		response["error"] = "昵称过短"
		c.JSON(400, response)
		return false
	}
	return true
}

func checkUserNotExist(username string, c *gin.Context) bool {
	response := ResponseInit()
	exist, err := model.DBIsUserExist(username)
	if err != nil {
		dbError(c)
		return false
	}

	if exist {
		response["error"] = "用户名已被注册"
		c.JSON(400, response)
		return false
	}

	return true
}

func fmtCheckUsername(username string) bool {
	if len(username) < 3 {
		return false
	}
	return true
}

func fmtCheckPassword(password string) bool {
	if len(password) < 6 {
		return false
	}
	return true
}

func fmtCheckTelephone(telephone string) bool {
	matched, _ := regexp.MatchString(`^((13[0-9])|(15[^4])|(18[0,2,3,5-9])|(17[0-8])|147)\d{8}$`, telephone)
	return matched
}

func fmtCheckEmail(email string) bool {
	matched, _ := regexp.MatchString(`^[a-zA-Z0-9_-]+@[a-zA-Z0-9_-]+(\.[a-zA-Z]+)+$`, email)
	return matched
}
