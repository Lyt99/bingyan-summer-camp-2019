package database

import (
	"database/sql"
	"fmt"
	"bytes"
	"shopsystem/model"
	"strings"

	//"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

//向数据库中添加用户
func CreateTableWithUser(db *sql.DB,username,password, tel, email string) {
	stmt, err := db.Prepare("INSERT systemuser SET username=?,password=?, tel=?, email=?")
	//id就不用添加了，数据库自己自动增加
	if err!=nil{
		fmt.Println(err,"有问题")
		return
	}
	_, err = stmt.Exec(username,password, tel, email)
	if err!=nil{
		fmt.Println(err,"有问题")
		return
	}
}

//看用户是否正确登录
func Checklogin(db *sql.DB,id int,password string)	int{
	row := db.QueryRow("SELECT username,password FROM userinfo WHERE id=?",id )
	//var uid int
	var get_username,get_password string
	/*
	   row：Scan()-->将查询的结果从row取出
	       err对象
	       判断err是否为空，
	           为空，查询有结果，数据可以成功取出
	           不为空，没有数据，sql: no rows in result set
	*/
	err := row.Scan(&get_username,&get_password)
	//fmt.Println(err)
	if err != nil {
		fmt.Println("查无数据。。",err)
		return 0
	} else {
		//password和get_password相同，也就是密码正确
		get_password = model.AesDecrypt(get_password)
		//先进行解密
		if  strings.Compare(get_password,password)==0{
			fmt.Println("欢迎您，",get_username)
			return 1
		}
		fmt.Println("密码错误，请重新登录")
		return 2
	}
}
//比较看用户是否已经注册
func Checksignup(db *sql.DB,username string)int{
	row := db.QueryRow("SELECT tel,email FROM userinfo WHERE username=?",username )
	//var uid int
	var tel,email string
	/*
	   row：Scan()-->将查询的结果从row取出
	       err对象
	       判断err是否为空，
	           为空，查询有结果，数据可以成功取出
	           不为空，没有数据，sql: no rows in result set
	*/
	err := row.Scan(&tel,&email )

	//fmt.Println(err)
	if err != nil {
		fmt.Println("查无数据。。",err)
		return 0
	} else {
		fmt.Println("该用户名可以进行注册")
		return 1
	}
}