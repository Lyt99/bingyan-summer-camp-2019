package database

import (
	"database/sql"
	"fmt"
	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"net/http"

	//"bytes"
	"shopsystem/model"
	"strings"

	//"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
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

//向数据库中添加用户
func CreateTableWithUser(c *gin.Context,username,nickname,password, tel, email string){
	stmt, err := db.Prepare("INSERT shopuser SET username=?, nickname=?,password=?, mobile=?, email=?")
	//id就不用添加了，数据库自己自动增加
	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"success":false,
			"error": "注册失败",
			"data":"",//失败的时候留空
		})
		return
	}
	_, err = stmt.Exec(username,nickname,password, tel, email)
	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"success":false,
			"error": "注册失败",
			"data":"",//失败的时候留空
		})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"success":true,
		"error": "",
		"data":"ok",	//不用传回用户的，而且没有id这一说
	})
	return
}
//比较看用户是否已经注册
func Checksignup(username string) (int,model.Userinfo){
	row := db.QueryRow("SELECT username,nickname,email,mobile FROM shopuser WHERE username=?",username )
	var user model.Userinfo
	/*
	   row：Scan()-->将查询的结果从row取出
	       err对象
	       判断err是否为空，
	           为空，查询有结果，数据可以成功取出
	           不为空，没有数据，sql: no rows in result set
	*/
	err := row.Scan(&user.Username,&user.Nickname,&user.Email,user.Mobile)
	if err != nil {
		fmt.Println("该用户名未被注册",err)
		return 0,user
	} else {
		fmt.Println("该用户名已经存在")
		user.Visittime++
		return 1,user
	}
}

//看用户是否正确登录
func Checklogin(c *gin.Context)(interface{}, error){
	var loginVals model.Register
	if err := c.ShouldBind(&loginVals); err != nil {
		return "", jwt.ErrMissingLoginValues
		//"incorrect Username or Password"
	}
	row := db.QueryRow("SELECT nickname,password,moblie,email FROM shopuser WHERE username=?",loginVals.Username )
	//var uid int
	var get_nickname,get_password,get_mobile,get_email string
	/*
	   row：Scan()-->将查询的结果从row取出
	       err对象
	       判断err是否为空，
	           为空，查询有结果，数据可以成功取出
	           不为空，没有数据，sql: no rows in result set
	*/
	err := row.Scan(&get_nickname,&get_password,&get_mobile,&get_email)
	//fmt.Println(err)
	if err != nil {
		c.String(http.StatusBadRequest,"用户名不存在")
		return "", jwt.ErrMissingLoginValues
	} else {
		//password和get_password相同，也就是密码正确
		get_password = model.AesDecrypt(get_password)
		//先进行解密
		if  strings.Compare(get_password,loginVals.Password)==0{
			fmt.Println("欢迎您，",get_nickname)
			return &model.Userinfo{
				Username: loginVals.Username,
				//Nickname:  get_nickname,
				//Mobile:  get_mobile,
				//Email: get_email,
			}, nil
		}
		fmt.Println("密码错误，请重新登录")
		return "账号或密码错误",jwt.ErrFailedAuthentication
	}
}
