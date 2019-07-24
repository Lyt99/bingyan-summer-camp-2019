package controller

import (
	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"onlineShopping/model"
	"strconv"
)

func GetPersonalInfo(c *gin.Context) {
	identity, _ := c.Get("username")
	username, _ := identity.(string)
	response := ResponseInit()
	if username == "" {
		return
	}
	user, err := model.DBSearchUser(username)
	if err != nil {
		dbError(c)
		return
	}

	response["success"] = true
	response["error"] = ""
	response["data"] = user
	c.JSON(200, response)
}

func UpdatePersonalInfo(c *gin.Context) {
	response := ResponseInit()
	identity, exist := c.Get("username")
	if !exist {
		response["error"] = "需要认证"
		c.JSON(401, response)
		return
	}
	username, _ := identity.(string)
	var json model.UserUpdateJSON
	if err := c.ShouldBindJSON(&json); err != nil {
		response["error"] = "更新信息不完整"
		c.JSON(400, response)
		return
	}
	if json.Password != "" {
		if updatePassword(username, json.Password, c) && updatePersonalInfo(json, username, c) {
			response["success"] = true
			response["error"] = ""
			response["data"] = "ok"
			c.JSON(200, response)
			return
		}
	} else {
		if updatePersonalInfo(json, username, c) {
			response["success"] = true
			response["error"] = ""
			response["data"] = "ok"
			c.JSON(200, response)
			return
		}
	}
}

func GetMyCommodities(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	identity, _ := claims["id"].(float64)
	userID := int(identity)
	response := ResponseInit()
	commodities, err := model.DBSearchUserCommodities(userID)
	if err != nil {
		dbError(c)
		return
	}
	response["success"] = true
	response["data"] = commodities
	response["error"] = nil
	c.JSON(200, response)
}

func GetUserInfo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	claims := jwt.ExtractClaims(c)
	fcurrentID, _ := claims["id"].(float64)
	currentID := int(fcurrentID)

	if currentID == id {
		GetPersonalInfo(c)
		return
	}
	response := ResponseInit()
	if err != nil {
		response["error"] = "用户ID格式错误"
		c.JSON(400, response)
		return
	}
	user, err := model.DBSearchUserByID(id)
	if err != nil {
		dbError(c)
		return
	}
	data := make(map[string]interface{})
	data["nickname"] = user.Username
	data["email"] = user.Email
	data["total_view_count"] = user.TotalViewCount
	data["total_collect_count"] = user.TotalCollectCount
	response["success"] = true
	response["error"] = ""
	response["data"] = data
	c.JSON(200, response)
}

func GetMyMessages(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	identity := claims["id"].(float64)
	userID := int(identity)
	response := ResponseInit()
	messages, err := model.DBGetUserMessages(userID)
	if err != nil {
		dbError(c)
		return
	}
	response["success"] = true
	response["error"] = ""
	response["data"] = messages
	c.JSON(200, response)
}

func updatePassword(username, password string, c *gin.Context) bool {
	response := ResponseInit()

	if !fmtCheckPassword(password) {
		response["error"] = "密码长度不足"
		c.JSON(400, response)
		return false
	}

	if err := model.DBUpdatePassword(username, password); err != nil {
		dbError(c)
		return false
	}
	return true
}

func updatePersonalInfo(json model.UserUpdateJSON, username string, c *gin.Context) bool {
	newJSON := jsonConvert(json, username)

	if fmtCheck(newJSON, c) {
		if err := model.DBUpdatePersonalInfo(newJSON); err != nil {
			dbError(c)
			return false
		}
		return true
	}
	return false
}

func jsonConvert(jsonUpdate model.UserUpdateJSON, username string) model.RegisterJSON {
	newJSON := model.RegisterJSON{username, jsonUpdate.Nickname, "default", jsonUpdate.Mobile, jsonUpdate.Email}
	return newJSON
}
