package util

import (
	"final_task/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func GetNameFromJWT(c echo.Context) string {
	user := c.Get("username").(*jwt.Token)
	claims := user.Claims.(*service.JwtCustomClaims)
	name := claims.UserName
	return name
}
