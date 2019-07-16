package controller

import (
	"demo/model"
	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"
)

// auth fail
func UnAuthFunc(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{"code": code, "message": message})
}

// fill in the payload of jwt
func FillPayload(data interface{}) jwt.MapClaims {
	if user, ok := data.(*model.User); ok {
		return jwt.MapClaims{"username": user.Username, "authority": user.Authority}
	}
	return jwt.MapClaims{}
}

//Create this middle ware
func MiddleWareInit() *jwt.GinJWTMiddleware {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:         "You have to login!",
		Key:           []byte("secret key salt"),
		Timeout:       time.Hour,
		MaxRefresh:    time.Hour,
		IdentityKey:   "username",
		PayloadFunc:   FillPayload,
		Authenticator: Login, // authentication
		Unauthorized:  UnAuthFunc,
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	return authMiddleware
}


func Login(c *gin.Context) (interface{}, error) { //authCallback?
	username := c.PostForm("username")
	password := c.PostForm("password")

	result := model.DB_search_user(username)
	if user, ok := result.(*model.User); ok {
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"info": "wrong password!"})
		}
		c.JSON(http.StatusOK, gin.H{
			"info":      "login succeed!",
			"ID":        user.ID,
			"username":  user.Username,
			"password":  user.Password,
			"authority": user.Authority,
		})
		return user, nil
	} else {
		c.JSON(http.StatusOK, gin.H{"info": "no such user!"})
	}
	return nil, jwt.ErrFailedAuthentication
}
