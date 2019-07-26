package controller

import (
	"final_task/model"
	"final_task/serializer"
	"final_task/service"
	"final_task/util"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

// GET /me 获取个人信息
func MeGetInfo(c echo.Context) (e error) {
	name := util.GetNameFromJWT(c)
	u := model.User{
		UserName: name,
	}
	if err := u.FindUserAsName(); err != nil {
		return c.JSON(http.StatusConflict, serializer.Response{
			Success: false,
			Error:   "这一定是删了自己账号然后又拿来访问才会出现的问题，我有时间给JWT加个过期，但是说实话这也没有删账号这种操作啊",
			Data:    nil,
		})
	} else {
		var r serializer.UserInfoResponse
		return c.JSON(http.StatusOK, r.SuccessResponse(u))
	}

}

// POST /me 更新个人信息
func MePostInfo(c echo.Context) (e error) {
	name := util.GetNameFromJWT(c)
	var s service.UserUpdateService
	return c.JSON(http.StatusOK, s.Updater(name))

}

// GET /me/commodities 获取个人发布商品
func MeGetCommodities(c echo.Context) (e error) {
	name := util.GetNameFromJWT(c)
	var s service.MyCommodityListService
	s.Username = name
	return c.JSON(http.StatusOK, s.GetMyCommodityList())
}

// GET /me/collections 获取我收藏的商品
func MeGetCollections(c echo.Context) (e error) {
	name := util.GetNameFromJWT(c)
	var s service.MyCollectionsGetService

	s.Username = name
	errResponse, mBL := s.GetMyCollections()
	if errResponse != nil {
		return c.JSON(http.StatusServiceUnavailable, errResponse)
	}
	return c.JSON(http.StatusOK, serializer.Response{
		Success: true,
		Error:   "",
		Data:    mBL.CBL,
	})

}

// POST /me/collections 收藏某件商品
func MePostCollections(c echo.Context) (e error) {
	name := util.GetNameFromJWT(c)
	var s service.CollectionRegisterService
	if s.ID, e = strconv.Atoi(c.Param("id")); e != nil {
		return c.JSON(http.StatusBadRequest, serializer.Response{
			Success: false,
			Error:   "商品ID格式不正确",
			Data:    nil,
		})
	}
	errResponse := s.Register(name)
	if errResponse != nil {
		return c.JSON(http.StatusNotFound, errResponse)
	}
	return c.JSON(http.StatusOK, serializer.Response{
		Success: true,
		Error:   "",
		Data:    "ok",
	})
}

// DELETE /me/collections 删除某件商品
func MeDeleteMeCollections(c echo.Context) (e error) {
	name := util.GetNameFromJWT(c)
	var s service.CollectionDeleteService

	s.ID, e = strconv.Atoi(c.FormValue("id"))
	if e != nil {
		return c.JSON(http.StatusBadRequest, serializer.Response{
			Success: false,
			Error:   "商品ID格式不正确",
			Data:    nil,
		})
	}
	s.Username = name
	if errResponse := s.DeleteCollection(); errResponse != nil {
		return c.JSON(http.StatusServiceUnavailable, errResponse)
	}
	return c.JSON(http.StatusOK, serializer.Response{
		Success: true,
		Error:   "",
		Data:    "ok",
	})
}
