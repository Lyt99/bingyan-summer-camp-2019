package middleware

import (
	"shopsystem/model"
	"shopsystem/database"
	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func MiddleWareInit () *jwt.GinJWTMiddleware  {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm: "You have to login!",
		Key:[]byte("secret key salt"),
		Timeout:time.Hour,
		MaxRefresh:time.Hour,
		IdentityKey: "id",
		PayloadFunc:PayloadFunc,

		//用来判断用户是否是正确用户
		Authenticator: database.Checklogin, // authentication
		//如果认证不成功的处理
		Unauthorized:UnauthorizedFun,
		TokenLookup:"header: Authorization, query: token, cookie: jwt",
		TokenHeadName:"Bearer",
		TimeFunc:time.Now,
	})
	if err != nil{
		log.Fatal("JWT Error:"+err.Error())
		return nil
	}
	return authMiddleware
}

//put user id into token
func PayloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(*model.Userinfo); ok {
		return jwt.MapClaims{
			"id": v.Username,					//这里“id”就是代表着 IdentityKey
		}
	}
	return jwt.MapClaims{}
}

//return login failed message
func UnauthorizedFun(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}

func HelloHandler(c *gin.Context) {
	//func ExtractClaims(c *gin.Context) jwt.MapClaims
	//用来将 Context 中的数据解析出来赋值给 claims
	//其实是解析出来了 JWT_PAYLOAD 里的内容
	//claims := jwt.ExtractClaims(c)

	id, err := c.Get("id")
	if !err {
		log.Println("warning: id get failed")
	}
	c.JSON(200, gin.H{
		"userID":   id,
		"text":     "welcome",
	})
}