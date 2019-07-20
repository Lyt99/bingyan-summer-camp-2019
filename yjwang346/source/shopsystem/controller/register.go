package controller

//这里是进行注册相关事情的
import (
	"github.com/gin-gonic/gin"
	"net/http"

	"shopsystem/database"
	"shopsystem/model"
)

func Signup(c *gin.Context){
	var newuser model.Register
	//newuser.UserName = c.PostForm("username")
	if err := c.ShouldBindJSON(&newuser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success":false,
			"error": err.Error(),
			"data":"",//失败的时候留空
		})

		return
	}
	password := model.AesEncrypt(newuser.Password)//进行加密,密码换成了进行加密之后得到的密码

	temp,_ := database.Checksignup(newuser.Username)
	if temp!=0{
		//c.String(http.StatusBadRequest,"注册失败，用户名重复，请更换用户名重新注册")		//这里是也要用那个response的格式吗？
		c.JSON(http.StatusBadRequest, gin.H{
			"success":false,
			"error": "用户名已存在",
			"data":"",//失败的时候留空
		})
		return
	}
	//加入数据到数据库中
	database.CreateTableWithUser(c,newuser.Username,newuser.Nickname,password,newuser.Mobile, newuser.Email)
	return
}