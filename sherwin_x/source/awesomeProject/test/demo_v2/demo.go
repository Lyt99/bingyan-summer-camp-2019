package main

import (
	"context"
	"fmt"
	"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type msg struct {
	id      string
	psw     string
	name    string
	tel     string
	email   string
	item    string
	context string
}

var identityKey = "id"
//var db *mongo.Database
var userColl *mongo.Collection
var ctx context.Context

//init MongoDB
func init() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
		fmt.Println("init failed")
	}
	//db = client.Database("demo")
	userColl = client.Database("demo").Collection("user")
}

//authority
func userCallback(c *gin.Context) (interface{}, error) {
	user := msg{}
	user.id = c.PostForm("ID")
	user.psw = c.PostForm("password")
	result := userColl.FindOne(ctx, bson.M{
		"ID":  user.id,
		"psw": user.psw})
	if err := result.Decode(&user); err != nil {
		fmt.Println("user login failed")
		return nil, jwt.ErrFailedAuthentication
	} else {
		fmt.Println("user been login")
		c.JSON(200, gin.H{"state": "success"})
		return &msg{
			name:  user.name,
			tel:   user.tel,
			email: user.email,}, nil
	}
}
func adminCallback(c *gin.Context) (interface{}, error) {
	user := msg{}
	user.id = c.PostForm("ID")
	user.psw = c.PostForm("password")
	if user.id == "admin" && user.psw == "12138" {
		return "200", nil
	} else {
		return nil, jwt.ErrFailedAuthentication
	}
}

/*
func authPrivCallback(data interface{}, c *gin.Context) bool {
	user:=msg{}
	user.id=c.PostForm("ID")
	user.psw=c.PostForm("password")
	if user.id=="admin"&&user.psw=="12138"{
		return true
	}else {return false}
	}*/

func main() {
	r := gin.New()

	//Middleware
	userMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{

		Realm:      "test",
		Key:        []byte("sherwin"),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		//put user id into token
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*msg); ok {
				return jwt.MapClaims{
					identityKey: v.id,
				}
			}
			return jwt.MapClaims{}
		},
		Authenticator: userCallback,
		//Authorizator: adminPrivCallback,
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		IdentityKey:   "id",
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})
	adminMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:         "test",
		Key:           []byte("sherwin"),
		Timeout:       time.Hour,
		MaxRefresh:    time.Hour,
		Authenticator: adminCallback,
		//Authorizator: adminPrivCallback,
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	//Router
	r.POST("/sign", SignHander)

	user := r.Group("/user")
	user.POST("/login", userMiddleware.LoginHandler)
	user.Use(userMiddleware.MiddlewareFunc())
	{
		user.POST("/update", UpdateHandler)
	}

	admin := r.Group("/admin")
	admin.POST("/login", adminMiddleware.LoginHandler)
	admin.Use(adminMiddleware.MiddlewareFunc())
	{
		admin.GET("/hello", helloHandler)
		admin.POST("/find", findHandler)
		admin.POST("/show", showHandler)
		admin.POST("/delete", delHandler)
	}

	_ = r.Run(":8080")
}

func SignHander(c *gin.Context) {
	user := msg{}
	user.id = c.PostForm("ID")
	user.psw = c.PostForm("password")
	user.name = c.PostForm("name")
	user.tel = c.PostForm("TEL")
	user.email = c.PostForm("e-mail")
	fmt.Println("submiting...")
	AddUser(user)
	c.JSON(200, gin.H{"html": "success"})
	return
}
func AddUser(newUser msg) {
	if result, err := userColl.InsertOne(ctx, bson.M{
		"ID":    newUser.id,
		"psw":   newUser.psw,
		"name":  newUser.name,
		"tel":   newUser.tel,
		"email": newUser.email}); err == nil {
		log.Println(result)
		fmt.Println("submit success！")

	} else {
		fmt.Print("submit failed")
	}
}

func UpdateHandler(c *gin.Context) {
	id, err := c.Get(identityKey)
	if !err {
		log.Println("id_get_failed")
		fmt.Println(id)
	}
	//fmt.Println(id)
	newdate := msg{}
	newdate.item = c.PostForm("item")
	newdate.context = c.PostForm("context")
	if result, err := userColl.UpdateOne(
		ctx, bson.M{"ID": id},
		bson.M{"$set": bson.M{newdate.item: newdate.context}}); err == nil {
		log.Println(result)
		log.Println(id)
	} else {
		log.Fatal(err)
	}
}
func helloHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"ID":   "shrwin",
		"text": "welcome",
	})
}
func delHandler(c *gin.Context) {
	del := msg{}
	del.id = c.PostForm("ID")
	if result, err := userColl.DeleteOne(ctx, bson.M{"ID": del.id}); err == nil {
		log.Println(result)
	} else {
		log.Println("delete failed")
	}
}
func findHandler(c *gin.Context) {
	find := msg{}
	find.id = c.PostForm("ID")
	result := userColl.FindOne(ctx, bson.M{"ID": find.id})
	if err := result.Decode(&find); err != nil {
		log.Println("not found")
	}
	log.Println(result)
	c.JSON(200, result)
}
func showHandler(c *gin.Context) {
	show := msg{}
	show.id = c.PostForm("ID")

}
