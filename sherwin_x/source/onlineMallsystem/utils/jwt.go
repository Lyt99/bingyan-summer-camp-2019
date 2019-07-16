package utils

import (
	"crypto/md5"
	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"onlineMallsystem/conf/Err"
	"onlineMallsystem/conf/msg"
	_ "onlineMallsystem/controller"
	"onlineMallsystem/model"
	"time"
)

//get user's token
func GetToken() *jwt.GinJWTMiddleware {
	Taken, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:         "test",
		Key:           []byte("sherwin"),
		Timeout:       time.Hour,
		MaxRefresh:    time.Hour,
		PayloadFunc:   payload,
		Authenticator: userCallback,
		Unauthorized:  unauthorized,
		LoginResponse: loginResponse,
		IdentityKey:   "id",
		TimeFunc:      time.Now,
	})
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
		return nil
	} else {
		return Taken
	}
}

//user authority
func userCallback(c *gin.Context) (interface{}, error) {
	log.Println(">>>User Authoring<<<")
	user := msg.LoginForm{}

	//bind login massage
	if err := c.ShouldBind(&user); err != nil {
		//c.JSON(200, Err.BindingFailed)
		return nil, Err.BindingFailed
	}

	//check sign
	filter := bson.M{
		"username": user.Username}
	if _, err := model.FindUser(filter); err != nil {
		//c.JSON(200, Err.UserNotExist)
		return nil, Err.UserNotExist
	}

	//encode psw to md5
	pswMd5 := md5.New()
	pswMd5.Write([]byte(user.Psw))
	user.Psw = string(pswMd5.Sum(nil))

	//find user in db
	filter = bson.M{
		"username": user.Username,
		"psw":      user.Psw}
	if res, err := model.FindUser(filter); err != nil {
		//c.JSON(200, Err.WrongPsw)
		return nil, Err.WrongPsw
	} else {
		user.Id = res.Id
	}
	return &msg.LoginForm{Id: user.Id,}, nil
}

//put user id into token
func payload(data interface{}) jwt.MapClaims {
	if v, ok := data.(*msg.LoginForm); ok {
		return jwt.MapClaims{
			msg.IdentityKey: v.Id,
		}
	}
	return jwt.MapClaims{}
}

//return login failed message
func unauthorized(c *gin.Context, code int, message string) {
	c.JSON(200, gin.H{
		"success": false,
		"error":   message,
		"data":    "",
	})
}

//return login success message
func loginResponse(c *gin.Context, code int, token string, time time.Time) {
	c.JSON(200, gin.H{
		"success": true,
		"error":   "",
		"data":    token,
	})
}

//auth test
func HelloHandler(c *gin.Context) {
	log.Println(">>>User Auth Test<<<")
	if Id, err := c.Get("id"); !err {
		c.JSON(200, gin.H{"state": "wrong!", "_id": Id})
	} else {
		c.JSON(200, gin.H{"state": "welcome!", "_id": Id})
	}
}
