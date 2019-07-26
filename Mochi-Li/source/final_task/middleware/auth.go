package middleware

import (
	"final_task/serializer"
	"final_task/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"net/http"
)

// 获取登陆用户
func CurrentUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		name := util.GetNameFromJWT(c)
		c.Set("username", name)
		return next(c)
	}
}

// 需要登陆
func AuthRequired(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		if user, _ := c.Get("username").(*jwt.Token); user != nil {
			return next(c)

		} else {
			return c.JSON(http.StatusUnauthorized, serializer.Response{
				Success: false,
				Error:   "尚未登陆",
				Data:    nil,
			})
		}

	}
}
