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
	stmt, err := db.Prepare("INSERT shopuser SET username=?, nickname=?,password=?, mobile=?, email=?,total_view_count=?,total_collect_count=?")
	//id就不用添加了，数据库自己自动增加
	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"success":false,
			"error": "注册失败asdasd",
			"data":"",//失败的时候留空
		})
		return
	}
	zero:=0
	_, err = stmt.Exec(username,nickname,password, tel, email,zero,zero)
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
	var user model.Userinfo
	row,err := db.Query("SELECT * FROM shopuser WHERE username=?",username )
	if err!=nil{
		fmt.Println("query这里出错",err)
		return 0,user
	}
	/*
	   row：Scan()-->将查询的结果从row取出
	       err对象
	       判断err是否为空，
	           为空，查询有结果，数据可以成功取出
	           不为空，没有数据，sql: no rows in result set
	*/
	//err := row.Scan(&user.Username,&user.Password,&user.Nickname,user.Mobile,&user.Email,&user.Visittime,&user.Collcectcount)
	for row.Next()  {
		err = row.Scan(&user.Username,&user.Password,&user.Nickname,&user.Mobile,&user.Email,&user.Visittime,&user.Collcectcount)//
	}
	//这里是没有办法的办法，程序在这里没有加  ||或 的话，总是view和collect项有值，但是总是01，
	// 其他的却总是空字符串！
	if err != nil || user.Username==""||user.Password=="" {
		fmt.Println("该用户名未被注册",err)
		return 0,user
	}
	fmt.Println("user.Username\n",user.Username,"user.Password\n",user.Password)
	fmt.Println("user.Nickname\n",user.Nickname,"user.Mobile\n",user.Mobile,user.Email,user.Visittime,user.Collcectcount)
	fmt.Println("该用户名已经存在")
	return 1,user
}

/*
func Checksignup(username string) (int,model.Userinfo){
	row := db.QueryRow("SELECT username,nickname,mobile,email FROM shopuser WHERE username=?",username )
	var user model.Userinfo

	   //row：Scan()-->将查询的结果从row取出
	   //    err对象
	   //    判断err是否为空，
	   //        为空，查询有结果，数据可以成功取出
	   //        不为空，没有数据，sql: no rows in result set
	   //
	err := row.Scan(&user.Username,&user.Nickname,user.Mobile,&user.Email)
	if err != nil {
		fmt.Println("该用户名未被注册",err)
		return 0,user
	} else {
		fmt.Println("该用户名已经存在")

		return 1,user
	}
}
*/

func Checklogin2(c *gin.Context){
	Checklogin(c)
}
//看用户是否正确登录
func Checklogin(c *gin.Context)(interface{}, error){
	var loginVals model.Login
	if err := c.ShouldBind(&loginVals); err != nil {
		return "", jwt.ErrMissingLoginValues
		//"incorrect Username or Password"
	}
	//c.String(http.StatusOK,loginVals.Username)
	//row := db.QueryRow("SELECT nickname,password,moblie,email FROM shopuser WHERE username=?",loginVals.Username )
	row,err := db.Query("SELECT * FROM shopuser WHERE username=?",loginVals.Username )
	if err!=nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success":false,
			"error": "获得商品信息出现错误",
			"data":"",//失败的时候留空
		})
		return "错误",nil
	}
	var view,collect int
	var get_nickname,get_password,get_mobile,get_email,get_username string
	/*
	   row：Scan()-->将查询的结果从row取出
	       err对象
	       判断err是否为空，
	           为空，查询有结果，数据可以成功取出
	           不为空，没有数据，sql: no rows in result set
	*/

	for  row.Next(){
		err = row.Scan(&get_username,&get_password,&get_nickname,&get_mobile,&get_email,&view,&collect)
	}
	//fmt.Println(err)
	if err != nil || get_nickname==""||get_password=="" {
		c.String(http.StatusBadRequest,"用户名不存在")
		fmt.Println(err)
		return "", jwt.ErrMissingLoginValues
	} else {
		//password和get_password相同，也就是密码正确
		get_password = model.AesDecrypt(get_password)
		//先进行解密
		if  strings.Compare(get_password,loginVals.Password)==0{
			fmt.Println("欢迎您，",get_nickname)
			return &model.Userinfo{
				Username: loginVals.Username,
				Nickname:  get_nickname,
				Mobile:  get_mobile,
				Email: get_email,
			}, nil
		}
		fmt.Println("密码错误，请重新登录")
		return "账号或密码错误",jwt.ErrFailedAuthentication
	}
}
//更改用户的信息
func Changeinfo(c *gin.Context,username string) int {
	var newinfo model.Changeuserinfo
	if err := c.ShouldBind(&newinfo); err != nil {
		return 1
		//"incorrect Username or Password"
	}
	if newinfo.Password!=""{
		stmt, err := db.Prepare("UPDATE  shopuser SET password=? WHERE username=?")
		if err != nil {
			return 1
		}
		//执行修改操作
		_, err = stmt.Exec(newinfo.Password,username)
		if err != nil {
			return 1
		}
	}
	if newinfo.Nickname!=""{
		stmt, err := db.Prepare("UPDATE  shopuser SET nickname=? WHERE username=?")
		if err != nil {
			return 1
		}
		//执行修改操作
		_, err = stmt.Exec(newinfo.Nickname,username)
		if err != nil {
			return 1
		}
	}
	if newinfo.Email!=""{
		stmt, err := db.Prepare("UPDATE  shopuser SET email=? WHERE username=?")
		if err != nil {
			return 1
		}
		//执行修改操作
		_, err = stmt.Exec(newinfo.Email,username)
		if err != nil {
			return 1
		}
	}
	if newinfo.Mobile!=""{
		stmt, err := db.Prepare("UPDATE  shopuser SET mobile=? WHERE username=?")
		if err != nil {
			return 1
		}
		//执行修改操作
		_, err = stmt.Exec(newinfo.Mobile,username)
		if err != nil {
			return 1
		}
	}

	if newinfo.Mobile==""&&newinfo.Email==""&&newinfo.Nickname==""&&newinfo.Password=="" {
		c.JSON(http.StatusBadRequest,gin.H{
			"success":false,
			"error": "更改内容不得为空",
			"data":"",//失败的时候留空
		})
		return 1
	}
	return 0
}

//将增加过后的浏览数据写入数据库中
func Addvisttime(c *gin.Context,username string,visittime int){
	stmt, err := db.Prepare("UPDATE  shopuser SET total_view_count=? WHERE username=?")
	if err != nil {
		return
	}
	//执行修改操作
	_, err = stmt.Exec(visittime,username)
	if err!=nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"success":false,
			"error": "",
			"data":"以增加该用户的浏览数",//失败的时候留空
		})
		return
	}

}