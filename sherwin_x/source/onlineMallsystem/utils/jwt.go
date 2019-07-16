package utils

import (
	"crypto/md5"
	"errors"
	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"onlineMallsystem/conf"
	_ "onlineMallsystem/controller"
	"onlineMallsystem/model"
	"time"
)

var IdentityKey="id"

//get user's token
func GetToken(key string) *jwt.GinJWTMiddleware {
	Taken, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:         "test",
		Key:           []byte(key),
		Timeout:       time.Hour,
		MaxRefresh:    time.Hour,
		PayloadFunc:   PayloadFunc,
		Authenticator: UserCallback,
		Unauthorized:  UnauthorizedFunc,
		IdentityKey:   "id",
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
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
func UserCallback(c *gin.Context) (interface{}, error) {
	log.Println(">>>User Authoring<<<")
	user := conf.LoginForm{}

	//bind login massage
	if err := c.ShouldBind(&user); err != nil {
		return nil, errors.New("invalid message")
	}

	//check sign
	filter := bson.M{
		"type": user.UserType,
		"tel":  user.Tel,}
	if _, err := model.FindUser(filter); err != nil {
		return nil, errors.New("incorrect UserType or Tel")
	}

	//encode psw to md5
	pswMd5 := md5.New()
	pswMd5.Write([]byte(user.Psw))
	user.Psw = string(pswMd5.Sum(nil))

	//find user in db
	filter = bson.M{
		"type": user.UserType,
		"tel":  user.Tel,
		"psw":  user.Psw}
	if res, err := model.FindUser(filter); err != nil {
		return nil, errors.New("incorrect Password")
	}else{
		return res, nil
	}
}

//put user id into token
func PayloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(*conf.LoginForm); ok {
		return jwt.MapClaims{
			IdentityKey: v.Id,
		}
	}
	return jwt.MapClaims{}
}

//return login failed message
func UnauthorizedFunc(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}

//auth test
func HelloHandler(c *gin.Context) {
	log.Println(">>>User Auth Test<<<")
	if user,err:=c.Get("id");!err{
		c.JSON(200,gin.H{"state":"wrong!"})
	}else{
		c.JSON(200, gin.H{"user":user})
	}
}