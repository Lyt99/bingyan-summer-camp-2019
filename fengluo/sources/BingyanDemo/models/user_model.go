package models

import (
	"BingyanDemo/database"
	"fmt"
)

type User struct {
	Userid     string
	Password   string
	Username   string
	Useradmin 	bool
	Userphone	string
	Useremail	string
}

//--------------数据库操作-----------------

//插入
func InsertUser(user User) (int64, error) {
	return database.ModifyDB("insert into information(user_id,user_password,user_name,user_phone,user_email,user_admin) values (?,?,?,?,?,?)",
		user.Userid, user.Password, user.Username, user.Userphone, user.Useremail, user.Useradmin)
}
func FindRepeatId(con string) int {
	sql := fmt.Sprintf("select id from information %s", con)
	fmt.Println(sql)
	row := database.FindDB(sql)
	id := 0
	row.Scan(&id)
	return id
}

func IfNameRepeat(username string) int {
	sql := fmt.Sprintf("where user_name = '%s'", username)
	return FindRepeatId(sql)
}

func IfIdRepeat(userid string) int {
	sql :=fmt.Sprintf("where user_id = '%s'", userid)
	return FindRepeatId(sql)
}
func IfNil(userid string, username string, password string, userphone string, useremail string, ) int {
	if useremail == "" ||userid == "" ||userphone == "" ||username == ""||password == ""{
		return 1
	}
	return 0
}
func LoginCheck(username string,password string) int {
	sql :=fmt.Sprintf("where user_name = '%s' and user_password = '%s'", username,password)
	return FindRepeatId(sql)
}