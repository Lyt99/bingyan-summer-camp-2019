package help

//如果报错请注释掉改文件所有内容

import (
	//Package log implements a simple logging package. It defines a type, Logger, with methods for formatting output.
	//这个包主要用来格式化的输出日志
	//"log"
	//Package http provides HTTP client and server implementations.
	//Get, Head, Post, and PostForm make HTTP (or HTTPS) requests
	//这个包提供了 HTTP 客户端和服务端的实现
	//"net/http"
	//Package os provides a platform-independent interface to operating system functionality
	//这个包主要用来实现一个跨平台的接口来操作系统函数
	//"os"
	//Package time provides functionality for measuring and displaying time
	//The calendrical calculations always assume a Gregorian calendar, with no leap seconds
	//用来提供一些计算和显示时间的函数
	//一个 Gin 的 JWT 中间件
	"time" //JWT Middleware for Gin framework

	"github.com/appleboy/gin-jwt" //Package gin implements a HTTP web framework called gin
	//这个包用来实现一个 HTTP 的 web 框架
	"github.com/gin-gonic/gin"
)

//定义一函数用来处理请求
func helloHandler(c *gin.Context) {
	//func ExtractClaims(c *gin.Context) jwt.MapClaims
	//ExtractClaims help to extract the JWT claims
	//用来将 Context 中的数据解析出来赋值给 claims
	//其实是解析出来了 JWT_PAYLOAD 里的内容
	/*
		func ExtractClaims(c *gin.Context) jwt.MapClaims {
		claims, exists := c.Get("JWT_PAYLOAD")
		if !exists {
			return make(jwt.MapClaims)
		}

		return claims.(jwt.MapClaims)
		}
	*/
	claims := jwt.ExtractClaims(c)
	//func (c *Context) JSON(code int, obj interface{})
	//JSON serializes the given struct as JSON into the response body. It also sets the Content-Type as "application/json".
	//将内容序列化成 JSON 格式，然后放到响应 body 里，同时将 Content-Type 置为 "application/json"
	c.JSON(200, gin.H{
		"userID": claims["id"],
		"text":   "Hello World",
	})
}

//定义一个 User 的结构体
type User struct {
	UserID    string //有三个字段都是 string 类型
	FirstName string
	LastName  string
}

//定义一个回调函数，用来决断用户id和密码是否有效
func authCallback(c *gin.Context) (interface{}, error) {
	//这里的 admin 和 test 都以硬代码的方式写进来用来作判断
	//生产环境下使用数据库来存储账号信息，进行检验和判断

	return nil, nil
}

//定义一个回调函数，用来决断用户在认证成功的前提下，是否有权限对资源进行访问
func authPrivCallback(user interface{}, c *gin.Context) bool {
	if v, ok := user.(string); ok && v == "admin" {
		return true
	}
	return false
}

//定义一个函数用来处理，认证不成功的情况
func unAuthFunc(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}

//main 包必要一个 main 函数，作为起点
func main() {
	//func Default() *Engine
	//Default returns an Engine instance with the Logger and Recovery middleware already attached.
	//用来返回一个已经加载了Logger and Recovery中间件的引擎
	r := gin.Default()

	// the jwt middleware
	authMiddleware := &jwt.GinJWTMiddleware{
		//Realm name to display to the user. Required.
		//必要项，显示给用户看的域
		Realm: "zone name just for test",
		//Secret key used for signing. Required.
		//用来进行签名的密钥，就是加盐用的
		Key: []byte("secret key salt"),
		//Duration that a jwt token is valid. Optional, defaults to one hour
		//JWT 的有效时间，默认为一小时
		Timeout: time.Hour,
		// This field allows clients to refresh their token until MaxRefresh has passed.
		// Note that clients can refresh their token in the last moment of MaxRefresh.
		// This means that the maximum validity timespan for a token is MaxRefresh + Timeout.
		// Optional, defaults to 0 meaning not refreshable.
		//最长的刷新时间，用来给客户端自己刷新 token 用的
		MaxRefresh: time.Hour,
		// Callback function that should perform the authentication of the user based on userID and
		// password. Must return true on success, false on failure. Required.
		// Option return user data, if so, user data will be stored in Claim Array.
		//必要项, 这个函数用来判断 User 信息是否合法，如果合法就反馈 true，否则就是 false, 认证的逻辑就在这里
		Authenticator: authCallback,
		// Callback function that should perform the authorization of the authenticated user. Called
		// only after an authentication success. Must return true on success, false on failure.
		// Optional, default to success
		//可选项，用来在 Authenticator 认证成功的基础上进一步的检验用户是否有权限，默认为 success
		Authorizator: authPrivCallback,
		// User can define own Unauthorized func.
		//可以用来息定义如果认证不成功的的处理函数
		Unauthorized: unAuthFunc,
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		//这个变量定义了从请求中解析 token 的格式
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		//TokenHeadName 是一个头部信息中的字符串
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		//这个指定了提供当前时间的函数，也可以自定义
		TimeFunc: time.Now,
	}

	//func (mw *GinJWTMiddleware) LoginHandler(c *gin.Context)
	//LoginHandler can be used by clients to get a jwt token. Payload needs to be json in the form of {"username": "USERNAME", "password": "PASSWORD"}. Reply will be of the form {"token": "TOKEN"}.
	//将 /login 交给 authMiddleware.LoginHandler 函数来处理
	r.POST("/login", authMiddleware.LoginHandler)

	//func (group *RouterGroup) Group(relativePath string, handlers ...HandlerFunc) *RouterGroup
	//Group creates a new router group. You should add all the routes that have common middlwares or the same path prefix. For example, all the routes that use a common middlware for authorization could be grouped
	//创建一个组 auth
	auth := r.Group("/auth")
	//func (mw *GinJWTMiddleware) MiddlewareFunc() gin.HandlerFunc
	//MiddlewareFunc makes GinJWTMiddleware implement the Middleware interface.
	//auth 组中使用 MiddlewareFunc 中间件
	auth.Use(authMiddleware.MiddlewareFunc())
	{
		//如果是 /auth 组下的 /hello 就交给 helloHandler 来处理
		auth.GET("/hello", helloHandler)
		//func (mw *GinJWTMiddleware) RefreshHandler(c *gin.Context)
		//RefreshHandler can be used to refresh a token. The token still needs to be valid on refresh. Shall be put under an endpoint that is using the GinJWTMiddleware. Reply will be of the form {"token": "TOKEN"}.
		//如果是 /auth 组下的 /refresh_token 就交给 RefreshHandler 来处理
		auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	}

	r.Run(":8080") //在 0.0.0.0:8080 上启监听
}
