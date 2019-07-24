package controller

import (
	"fmt"
	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"onlineShopping/model"
	"strconv"
)

func SearchCommodities(c *gin.Context) {
	response := ResponseInit()
	var json model.SearchJSON
	if err := c.ShouldBindJSON(&json); err != nil {
		fmt.Println(err)
		response["error"] = "搜索信息不全或格式错误"
		c.JSON(400, response)
		return
	}

	if fmtCheckSearch(json, c) {
		commodities, err := model.DBSearchCommodities(json)
		if err != nil {
			response["error"] = "无法连接数据库"
			c.JSON(500, response)
			return
		}
		response["success"] = true
		response["data"] = commodities
		response["error"] = ""
		c.JSON(200, response)
	}
}

func DetailedCommodity(c *gin.Context) {
	response := ResponseInit()
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response["error"] = "查询格式有误"
		c.JSON(400, response)
		return
	}

	commodity, exist, err := model.DBSearchCommodity(id)
	if err != nil {
		response["error"] = "无法连接数据库"
		c.JSON(500, response)
		return
	}
	if !exist {
		response["error"] = "无商品信息"
		c.JSON(404, response)
		return
	}
	response["success"] = true
	response["error"] = nil
	response["data"] = commodity //bug here: how to drop id?
	c.JSON(200, response)
}

func ReleaseCommodity(c *gin.Context) {
	identity, _ := c.Get("username")
	username, _ := identity.(string)
	response := ResponseInit()
	var json model.ProductJSON

	if err := c.ShouldBindJSON(&json); err != nil {
		response["error"] = "更新信息不完整"
		c.JSON(400, response)
		return
	}

	if releaseProduct(json, username, c) {
		response["success"] = true
		response["error"] = ""
		response["data"] = "ok"
		c.JSON(200, response)
	}
}

func DeleteCommodity(c *gin.Context) {
	response := ResponseInit()
	claims := jwt.ExtractClaims(c)
	identity := claims["id"]
	if fuserID, ok := identity.(float64); ok {
		userID := int(fuserID)
		commodityID, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			response["error"] = "商品ID格式错误"
			c.JSON(400, response)
			return
		}

		commodity, exist, err := model.DBSearchCommodity(commodityID)
		if err != nil {
			dbError(c)
			return
		}
		if !exist {
			response["error"] = "没有此商品"
			c.JSON(400, response)
			return
		}
		if commodity.PubUser != userID {
			response["error"] = "没有删除权限"
			c.JSON(401, response)
			return
		} else {
			if err := model.DBDeleteCommodity(commodityID); err != nil {
				response["error"] = "删除失败"
				c.JSON(500, response)
				return
			}
			model.OSDeletePic(commodity.Picture)
			response["success"] = true
			response["error"] = nil
			response["data"] = "ok"
			c.JSON(200, response)
		}
	}
}

func GetHotWords(c *gin.Context) {
	words, err := model.DBGetHotWords()
	if err != nil {
		dbError(c)
		return
	}
	response := ResponseInit()
	response["success"] = true
	response["error"] = ""
	response["data"] = words
	c.JSON(200, response)
}

//this func is used to check whether the
//format is valid.
func fmtCheckSearch(json model.SearchJSON, c *gin.Context) bool {
	response := ResponseInit()
	if json.PageNo < 0 {
		response["error"] = "页码错误"
		c.JSON(400, response)
		return false
	}
	if json.PageSize < 0 {
		response["error"] = "每页最大商品数错误"
		c.JSON(400, response)
		return false
	}
	if json.Category < 0 || json.Category > 9 {
		response["error"] = "商品类别错误"
		c.JSON(400, response)
		return false
	}
	return true
}

func releaseProduct(json model.ProductJSON, username string, c *gin.Context) bool {
	if !fmtCheckProduct(json, c) {
		return false
	}
	user, err := model.DBSearchUser(username)
	if err != nil {
		fmt.Println("search user bug")
		dbError(c)
		return false
	}
	json.PubUser = user.ID
	err = model.DBReleaseProduct(json)
	if err != nil {
		fmt.Println("db write bug")
		dbError(c)
		return false
	}
	return true
}

func fmtCheckProduct(json model.ProductJSON, c *gin.Context) bool {
	response := ResponseInit()
	if !fmtCheckTitle(json.Title) {
		response["error"] = "标题不可以为空"
		c.JSON(400, response)
		return false
	}

	if !fmtCheckCategory(json.Category) {
		response["error"] = "类别不符合规范"
		c.JSON(400, response)
		return false
	}
	return true
}

func fmtCheckTitle(title string) bool {
	if len(title) < 1 {
		return false
	}
	return true
}

func fmtCheckCategory(category int) bool {
	if category < 0 || category > 9 {
		return false
	}
	return true
}
