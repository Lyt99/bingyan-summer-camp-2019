package controller

//这里是进行注册相关事情的
import (
	"fmt"
	"net/http"
	"database/sql"
	"github.com/gin-gonic/gin"

	"shopsystem/model"
	"shopsystem/database"
)
var db *sql.DB
//不过话说什么时候进行关闭连接？
//在使用Postman访问的时候，json里面的数据名称为:id\username\password\telephone\email
func init(){
	var err error
	db,err=sql.Open("mysql","root:0@tcp(localhost:3306)/blogweb_gin?charset=utf8")
	if err!=nil{
		fmt.Println(err)
		return
	}
	if err!=nil{
		fmt.Println(err,"ie")
	}
	//panic(err)
}

func Signup(c *gin.Context){
	var newuser model.Register
	//newuser.UserName = c.PostForm("username")
	if err := c.ShouldBindJSON(&newuser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error错误": err.Error()})
		return
	}
	password := newuser.Password
	//password := model.AesEncrypt(newuser.Psw)//进行加密,密码换成了进行加密之后得到的密码
	//先进行用户名查重
	temp := database.Checksignup(db,newuser.UserName)
	if temp==0{
		//加入数据库内容？
		database.CreateTableWithUser(db,newuser.UserName,password,newuser.Telephone, newuser.Email)
	} else {
		c.String(http.StatusBadRequest,"注册失败，用户名重复，请更换用户名重新注册")
		return
	}
	c.String(http.StatusOK,"注册成功！")
}