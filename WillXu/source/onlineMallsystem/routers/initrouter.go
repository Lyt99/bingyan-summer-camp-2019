package routers

import (
	"github.com/gin-gonic/gin"
	"onlineMallsystem/controllers"
	"onlineMallsystem/utils"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.Use(utils.Cors())                                  // 允许使用跨域请求,全局中间件
	r.POST("/user", controllers.SignUp)                  //注册
	r.POST("/user/login", utils.GetToken().LoginHandler) //登录

	r.Use(utils.GetToken().MiddlewareFunc()) //jwt中间件
	{
		user := r.Group("/user")
		user.GET("/:id", controllers.ShowUser) //查看某位用户资料

		me := r.Group("/me")
		me.GET("/hello", utils.HelloHandler)                    //测试
		me.GET("", controllers.ShowMe)                          //查看个人资料
		me.POST("", controllers.UpdateMe)                       //修改个人资料
		me.GET("/commodities", controllers.MyCommodities)       //查看我的发布
		me.GET("/collections", controllers.MyCollections)       //查看我的收藏
		me.POST("/collections", controllers.NewCollection)      //收藏某个商品
		me.DELETE("/collections", controllers.DeleteCollection) //取消某个收藏

		commodities := r.Group("/commodities")
		commodities.GET("", controllers.GetCommodities)   //获取商品列表
		commodities.GET("/hot", controllers.GetHotSearch) //获取热搜词
		commodities.POST("", controllers.NewCommodity)    //发布新商品

		commodity := r.Group("/commodity")
		commodity.GET("/:id", controllers.DetailCommodity)    //某个商品详情
		commodity.DELETE("/:id", controllers.DeleteCommodity) //删除某个商品

		r.POST("/pics", controllers.UploadPic) //上传图片
	}
	return r
}
