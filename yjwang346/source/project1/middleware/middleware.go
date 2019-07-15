package middleware



import (

	"project1/controllers"
	"project1/models"
	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"log"
	"time"

)



func GetUserToken() *jwt.GinJWTMiddleware {

	userToken,err:=jwt.New(&jwt.GinJWTMiddleware{
		Realm:"test",
		Key:[]byte("user"),
		Timeout:time.Hour,
		MaxRefresh:time.Hour,

		//put user id into token
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*models.LoginForm); ok {
				return jwt.MapClaims{
					models.IdentityKey: v.Id,
				}
			}

			return jwt.MapClaims{}

		},

		Authenticator:controllers.UserCallback,

		//Authorizator: adminPrivCallback,

		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		IdentityKey: "id",
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc: time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
		return nil
	}else {return userToken}
}


func GetAdminToken() *jwt.GinJWTMiddleware{
	adminTaken, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:"test",
		Key:[]byte("sherwin"),
		Timeout:time.Hour,
		MaxRefresh:time.Hour,
		Authenticator:controllers.AdminCallback,

		//Authorizator: adminPrivCallback,

		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc: time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
		return nil
	}else {return adminTaken}
}