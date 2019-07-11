package model

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)


type User struct{
	ID int
	Username string
	Password string
	Phonenumber string
	Email string
	Authority int
}


func DB_register(username string, password string, phonenumber string, email string, authority int) {
	db, err := sql.Open("mysql", "root:@/test?charset=utf8")
	checkErr(err)

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil{
		fmt.Println(err)
	}

	hashed_password := string(hash)

	stmt, err := db.Prepare("INSERT userinfo SET username=?, password=?, phonenumber=?, email=?, authority=?")
	checkErr(err)

	_, err = stmt.Exec(username, hashed_password, phonenumber, email, authority)
	checkErr(err)

	db.Close()
}


func DB_search_user(username string) interface{} {
	db, err := sql.Open("mysql", "root:@/test?charset=utf8")
	checkErr(err)

	rows, err := db.Query("SELECT * FROM userinfo WHERE username=?",username)
	checkErr(err)

	db.Close()

	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.Username, &user.Password, &user.Phonenumber, &user.Email, &user.Authority)
		checkErr(err)
		return user
	}
	return nil
}


func DB_isExist(username string) bool {
	db, err := sql.Open("mysql", "root:@/test?charset=utf8")
	checkErr(err)

	rows, err := db.Query("SELECT * FROM userinfo WHERE username=?",username)
	checkErr(err)

	db.Close()

	return rows.Next()
}


func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}


func DB_delete_user(target_user User) {
	db, err := sql.Open("mysql", "root:@/test?charset=utf8")
	checkErr(err)

	stmt, err := db.Prepare("DELETE FROM userinfo WHERE username=?")
	checkErr(err)

	_, err = stmt.Exec(target_user.Username)
	db.Close()
}