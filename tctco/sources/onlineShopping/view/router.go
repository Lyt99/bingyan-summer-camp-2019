package view

import (
	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"onlineShopping/controller"
)

func RouterInit(authMiddleware *jwt.GinJWTMiddleware) *gin.Engine {
	r := gin.Default()
	loginInit(r, authMiddleware)
	freeAccessInit(r)
	staticFilesInit(r)
	meGroupInit(r, authMiddleware)
	picGroupInit(r, authMiddleware)
	commoditiesGroupInit(r, authMiddleware)
	commodityGroupInit(r, authMiddleware)
	userGroupInit(r, authMiddleware)
	return r
}

func staticFilesInit(r *gin.Engine) {
	r.StaticFS("/upload", http.Dir("./upload"))
}

func loginInit(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) {
	r.POST("/user/login", authMiddleware.LoginHandler)
}

//these routes don't need authentication
func freeAccessInit(r *gin.Engine) {
	r.GET("/", controller.Hello)
	r.POST("/user", controller.Register)
	r.GET("/commodities", controller.SearchCommodities)
	r.GET("/commodity/:id", controller.DetailedCommodity)
	r.GET("/commodities/hot", controller.GetHotWords)
}

//these routes are related to the user
func meGroupInit(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) {

	me := r.Group("/me")
	me.Use(authMiddleware.MiddlewareFunc())
	{
		me.GET("", controller.GetPersonalInfo)
		me.POST("", controller.UpdatePersonalInfo)
		me.GET("/messages", controller.GetMyMessages)
		me.GET("/commodities", controller.GetMyCommodities)
		me.GET("/collections", controller.GetMyCollections)
		me.POST("/collections", controller.AddToCollections)
		me.DELETE("/collections", controller.DeleteFromCollections)
	}
}

//this route is used for uploading pics
func picGroupInit(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) {
	pic := r.Group("/pic")
	pic.Use(authMiddleware.MiddlewareFunc())
	{
		pic.POST("", controller.UploadPic)
	}
}

//this route is used to release products
func commoditiesGroupInit(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) {
	commodities := r.Group("/commodities")
	commodities.Use(authMiddleware.MiddlewareFunc())
	{
		commodities.POST("", controller.ReleaseCommodity)
	}
}

//this route is used to delete commodities
func commodityGroupInit(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) {
	commodity := r.Group("/commodity")
	commodity.Use(authMiddleware.MiddlewareFunc())
	{
		commodity.DELETE("/:id", controller.DeleteCommodity)
	}
}

//this route is used to get other user's profile
func userGroupInit(r *gin.Engine, authMiddleware *jwt.GinJWTMiddleware) {
	user := r.Group("/user")
	user.Use(authMiddleware.MiddlewareFunc())
	{
		user.GET("/:id", controller.GetUserInfo)
	}
}