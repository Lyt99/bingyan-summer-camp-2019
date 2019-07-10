package gin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
*/

/*
func main() {
	// Engin
	router := gin.Default()
	//router := gin.New()

	router.GET("/hello", func(context *gin.Context) {
		log.Println(">>>> hello gin start <<<<")
		context.JSON(200, gin.H{
			"code":    200,
			"success": true,
		})
	})
	// 指定地址和端口号
	router.Run()
}
*/

func main() {

	gin.SetMode(gin.DebugMode) //全局设置环境，此为开发环境，线上环境为gin.ReleaseMode
	router := gin.Default()    //获得路由实例

	//添加中间件
	router.Use(Middleware)
	//注册接口
	router.GET("/get", GetHandler)
	router.POST("/post", PostHandler)
	router.PUT("/put", PutHandler)
	router.DELETE("/delete", DeleteHandler)
	//监听端口
	http.ListenAndServe(":8080", router)
}

func Middleware(c *gin.Context) {
	fmt.Println("this is a middleware!")
}

func GetHandler(c *gin.Context) {
	value, exist := c.GetQuery("key")
	if !exist {
		value = "the key is not exist!"
	}
	c.Data(http.StatusOK, "text/plain", []byte(fmt.Sprintf("get success! %s\n", value)))
	return
}
func PostHandler(c *gin.Context) {
	type JsonHolder struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	}
	holder := JsonHolder{Id: 1, Name: "my name"}
	//若返回json数据，可以直接使用gin封装好的JSON方法
	c.JSON(http.StatusOK, holder)
	return
}
func PutHandler(c *gin.Context) {
	c.Data(http.StatusOK, "text/plain", []byte("put success!\n"))
	return
}
func DeleteHandler(c *gin.Context) {
	c.Data(http.StatusOK, "text/plain", []byte("delete success!\n"))
	return
}