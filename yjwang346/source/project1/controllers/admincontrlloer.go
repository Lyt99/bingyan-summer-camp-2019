package controllers

import (
	"database/sql"
	"fmt"
	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"project1/models"

	_ "github.com/go-sql-driver/mysql"
)
//查询一条
func Getuserinfo(c *gin.Context) {
	/*
	   查询一条
	*/
	db, _ := sql.Open("mysql", "root:hanru1314@tcp(localhost:3306)/bolgweb_gin?charset=utf8")
	var s models.SignForm
	s.Id := c.Params("Id")
	row := db.QueryRow("SELECT Id,UserName,Tel,Email FROM userinfo WHERE Id=?", s.Id)
	var Id uint16
	var UserName string
	var Tel string
	//var Psw string
	var Email string
	/*
	   row：Scan()-->将查询的结果从row取出
	       err对象
	       判断err是否为空，
	           为空，查询有结果，数据可以成功取出
	           不为空，没有数据，sql: no rows in result set
	*/
	err := row.Scan(&Id, &UserName, &Tel, &Email)

	//fmt.Println(err)
	if err != nil {
		fmt.Println("查无数据。。")
	} else {
		fmt.Println(Id,UserName, Tel, Email)
	}
	db.Close()
}

//查询多条,但是这个方法比较麻烦
func getusersinfo() {
	/*
	   查询操作：
	*/
	//step2:打开数据库，建立连接
	db, _ := sql.Open("mysql", "root:0@tcp(localhost:3306)/bolgweb_gin?charset=utf8")

	//stpt3：查询数据库
	rows, err := db.Query("SELECT Id,UserName,Tel,Email FROM userinfo")
	if err != nil {
		fmt.Println("查询有误。。")
		return
	}
	//fmt.Println(rows.Columns()) //[uid username departname created]
	//创建slice，存入struct，
	datas := make([] models.SignForm, 0)
	//step4：操作结果集获取数据
	for rows.Next() {
		var Id uint16
		var UserName string
		var Tel string
		var Psw string
		var Email string
		if err := rows.Scan(&Id, &UserName, &Tel, &Psw,&Email); err != nil {
			fmt.Println("获取失败。。")
		}
		//每读取一行，创建一个user对象，存入datas2中
		user := models.SignForm{Id, Psw, UserName, Tel,Email}
		datas = append(datas, user)
	}
	//step5：关闭资源
	rows.Close()
	db.Close()

	for _, v := range datas {
		fmt.Println(v)
	}
}
/*
查询：处理查询后的结果：
    思路一：创建结构体
    思路二：将数据，存入到map中
*/

//将数据库里的user全部输出,这个功能应该没有问题
//没有使用到gin.Context;
func Check(err error){     //因为要多次检查错误，所以干脆自己建立一个函数。
	if err!=nil{
		fmt.Println(err)
	}
}
func GetUsersinfo(c *gin.Context){
	db,err:=sql.Open("mysql","root:123456@tcp(127.0.0.1:3306)/employee")
	Check(err)
	//query

	rows,err := db.Query("SELECT * FROM message")
	Check(err)

	c.String(http.StatusOK,"所有用户的信息如下：")
	for rows.Next(){
		var s models.SignForm
		err=rows.Scan(&s.Id,&s.UserName,&s.Psw,&s.Tel,&s.Email)
		Check(err)
		fmt.Println(s)
	}
	rows.Close()
}
//删除
func Deletuser(c *gin.Context){
	db,err:=sql.Open("mysql","root:123456@tcp(127.0.0.1:3306)/employee")
	Check(err)

	var s models.SignForm
	temp := c.Params("Id")
	s.Id = uint16(temp)
	results,err := db.Exec("delete from student where UserName = ?", s.Id)
	Check(err)
	fmt.Println(results.RowsAffected())
}