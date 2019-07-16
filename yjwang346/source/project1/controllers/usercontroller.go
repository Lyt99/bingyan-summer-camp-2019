package controllers

import (
	"fmt"
	"net/http"

	//"github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	//"log"
	"database/sql"
	"project1/models"
)
//var db *sql.DB		//通过查询，知道db是*sql.DB类型,本来是想用静态变量的
var idnum int = 0
/*func InitMysql() {
	fmt.Println("InitMysql....")
	_,err:=sql.Open("mysql","root:0@tcp(localhost:3306)/mytest?charset=utf8")
	if err !=nil{
		fmt.Println("连接失败。。")
		return
	}
}*/

func Signup(c *gin.Context){
	db,err:=sql.Open("mysql","root:0@tcp(localhost:3306)/mytest?charset=utf8")
	defer db.Close()
	Check(err)
	idnum = idnumber(idnum,db)
	idnum++				//计数完之后还要进行一次+
	//但问题是，程序关闭之后，id也会变成零该怎么办？
	var newuser models.SignForm
	newuser.UserName = c.PostForm("username")
	newuser.Psw = c.PostForm("password")
	newuser.Tel = c.PostForm("telephone")
	newuser.Email = c.PostForm("email")
	CreateTableWithUser(db,idnum ,newuser.UserName,newuser.Psw,newuser.Tel, newuser.Email)
	//加入数据库内容？
	checksignup(db,idnum)
	c.String(http.StatusOK,"注册成功！")
}
func CreateTableWithUser(db *sql.DB,id int,UserName,Psw, Tel, Email string) {
	stmt, err := db.Prepare("INSERT userinfo SET id=?,UserName=?,Psw=?, Tel=?, Email=?")
	Check(err)

	_, err = stmt.Exec(id,UserName,Psw, Tel, Email)
	Check(err)
}
//操作数据库		存在疑问！！！


//比较看用户是否已经注册
func checksignup(db *sql.DB,id int) {
		//查询一条

		//db, _ := sql.Open("mysql", "root:hanru1314@tcp(localhost:3306)/bolgweb_gin?charset=utf8")
		row := db.QueryRow("SELECT Id,UserName,Tel,Email FROM userinfo WHERE uid=?", id)
		//var uid int
		//var username, departname, created string
		var check models.SignForm
		 /*  row：Scan()-->将查询的结果从row取出
	 	      err对象,判断err是否为空，
		           为空，查询有结果，数据可以成功取出
	 	          不为空，没有数据，sql: no rows in result set
		*/
		err := row.Scan(&check.Id, &check.UserName)

		//fmt.Println(err)
		if err != nil {
			fmt.Println("该用户名可以使用")
		} else {
			fmt.Println("该用户名已经使用，请更换用户名")
		}
}

//方法很笨，	用来得到数据库中有多少数据

func idnumber(idnum int,db *sql.DB) int {
	rows,err := db.Query("SELECT * FROM message")
	Check(err)

	for rows.Next(){
		idnum++
	}
	return idnum
}