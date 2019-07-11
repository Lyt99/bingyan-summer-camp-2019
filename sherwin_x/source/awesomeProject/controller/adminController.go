package controller

import (
	"awesomeProject/model"
	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"log"
)
//admin authority
func AdminCallback(c *gin.Context)(interface{},error){
	log.Println(">>>Admin Authoring<<<")
	user:=model.LoginForm{}
	_ =c.ShouldBindUri(&user)
	if user.Id==2019&&user.Psw=="0711"{
	return "200", nil
	}else {
	return nil, jwt.ErrFailedAuthentication}
}

// auth test
func HelloAdminHandler(c *gin.Context)  {
	log.Println(">>>Admin Auth Test<<<")
	c.JSON(200,gin.H{
		"ID":"admin",
		"text":"Do What You Want!",
	})
}

//find one user
//TBD
func FindHandler(c *gin.Context)  {
	findId:=c.PostForm("id")
	log.Println(">>>Admin Find User: %s<<<",findId)
	result := UserColl.FindOne(ctx, bson.M{"ID": findId})
	if err := result.Decode(&findId); err != nil {
		log.Println("not found")
	}
	log.Println(result)
	c.JSON(200,result)
}

//show all user
//TBD
func ShowHandler(c *gin.Context){
	log.Println(">>>Admin Show All Users<<<")

}

//delete user
func DelHandler(c *gin.Context)  {
	delId:=c.PostForm("id")
	log.Println(">>>Admin Delete User: %s<<<",delId)
	if result, err := UserColl.DeleteOne(ctx, bson.M{"ID": delId}); err == nil {
		log.Println(result)
	} else {
		log.Println("delete failed")
	}
}