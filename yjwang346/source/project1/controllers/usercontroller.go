package controllers

import (
	"fmt"
	"github.com/appleboy/gin-jwt"
	"log"
	"database/sql"
	"project1/models"
)
var db *sql.DB		//通过查询，知道db是*sql.DB类型

func InitMysql() {
	fmt.Println("InitMysql....")
	db,err:=sql.Open("mysql","root:0@tcp(localhost:3306)/mytest?charset=utf8")
	if err !=nil{
		fmt.Println("连接失败。。")
		return
	}
}

func signup(){
	CreateTableWithUser()
	//加入数据库内容？
	checksignup()
	fmt.Println("注册成功！")
}
//创建用户表
func CreateTableWithUser() {
	sql := `CREATE TABLE IF NOT EXISTS users(
        Id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
        Name VARCHAR(64),
        Psw VARCHAR(64),
		Tel VARCHAR(64),
		Email VARCHAR(64)
        );`
	//status INT(4),
	//createtime INT(10)
	ModifyDB(sql)
}
//操作数据库		存在疑问！！！
func ModifyDB(sql string, args ...interface{}) (int64, error) {
	result, err := db.Exec(sql, args...)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return count, nil
}

//比较看用户是否已经注册
func checksignup() {
	/*
	   查询一条
	*/
	db, _ := sql.Open("mysql", "root:hanru1314@tcp(127.0.0.1:3306)/mytest?charset=utf8")

	row := db.QueryRow("SELECT uid,username,departname,created FROM userinfo WHERE uid=?", 1)
	//var uid int
	//var username, departname, created string
	var check models.SignForm
	/*
	   row：Scan()-->将查询的结果从row取出
	       err对象
	       判断err是否为空，
	           为空，查询有结果，数据可以成功取出
	           不为空，没有数据，sql: no rows in result set
	*/
	err := row.Scan(&check.Id, &check.Name)

	//fmt.Println(err)
	if err != nil {
		fmt.Println("该用户名可以使用")
	} else {
		fmt.Println("该用户名已经使用，请更换用户名")
	}
	db.Close()
}