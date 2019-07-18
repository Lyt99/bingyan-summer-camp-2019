package database

import (
	"database/sql"
	"fmt"
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
		fmt.Println("查无数据。。")
		return 0
	} else {
		fmt.Println(tel,email)
		return 1
	}

}