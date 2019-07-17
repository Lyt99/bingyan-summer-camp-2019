package model

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

// i wish these functions to be as simple as they can
// and leave all sorts of errors and bad requests in the
//controller layer

type User struct {
	ID          int
	Username    string
	Password    string
	Phonenumber string
	Email       string
	Authority   int
}

func dbInit() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@/test?charset=utf8")
	if err != nil {
		fmt.Println("cannot create a connection pool")
		return nil, err
	}
	return db, nil
}

func DB_register(username string, password string, phonenumber string, email string, authority int) bool {
	db, _ := dbInit()

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) // store the hashed password
	if err != nil {
		fmt.Println("cannot hash password")
		return false
	}

	hashed_password := string(hash)

	stmt, err := db.Prepare("INSERT userinfo SET username=?, password=?, phonenumber=?, email=?, authority=?")
	if err != nil {
		fmt.Println("cannot prepare")
		return false
	}

	_, err = stmt.Exec(username, hashed_password, phonenumber, email, authority)
	if err != nil {
		fmt.Println("cannot execute")
		return false
	}
	defer db.Close()
	return true
}

func DB_isExist(username string) (bool, error) {
	db, _ := dbInit()

	rows, err := db.Query("SELECT * FROM userinfo WHERE username=?", username)
	if err != nil {
		fmt.Println("cannot query")
		return false, err
	}

	db.Close()
	return rows.Next(), nil
}

func DB_delete_user(username string) bool {
	db, _ := dbInit()

	stmt, err := db.Prepare("DELETE FROM userinfo WHERE username=?")
	if err != nil {
		fmt.Println("cannot prepare")
		return false
	}

	_, err = stmt.Exec(username)
	defer db.Close()
	return true
}

func DB_search_user_info(pageno, pagesize int, username string) ([]User, int) {
	db, _ := dbInit()

	users := make([]User, 0)
	var cnt int
	if username != "" { // search all the users!
		rows, err := db.Query("SELECT id, username, phonenumber, email, authority FROM userinfo WHERE username=?", username)
		if err != nil {
			return nil, 0
		}
		defer rows.Close()

		for rows.Next() {
			var user User
			cnt++
			rows.Scan(&user.ID, &user.Username, &user.Phonenumber, &user.Email, &user.Authority)
			users = append(users, user)
		}
	} else {
		rows, err := db.Query("SELECT id, username, phonenumber, email, authority FROM userinfo LIMIT ?,?", (pageno-1)*pagesize, pagesize)
		if err != nil {
			return nil, 0
		}

		for rows.Next() { // search a particular user
			var user User
			cnt++
			rows.Scan(&user.ID, &user.Username, &user.Phonenumber, &user.Email, &user.Authority)
			users = append(users, user)
		}
		if err = rows.Err(); err != nil {
			return nil, 0
		}
		defer rows.Close()
	}
	return users, cnt
} // this function doesn't return password

func DB_search_user(username string) interface{} {
	db, _ := dbInit()

	row, err := db.Query("SELECT * FROM userinfo WHERE username=?", username)
	if err != nil {
		fmt.Println("cannot query")
		return nil
	}

	var user User
	for row.Next() {
		row.Scan(&user.ID, &user.Username, &user.Password, &user.Phonenumber, &user.Email, &user.Authority)
		return &user // this is a practice of pointer!
	}
	defer row.Close()
	return nil
} // this function returns all info

func DB_update_info(username, new_tel, new_email string) bool {
	db, _ := dbInit()

	stmt, err := db.Prepare("UPDATE userinfo SET phonenumber=?, email=? WHERE username=?")
	if err != nil {
		fmt.Println("cannot prepare")
		return false
	}

	_, err = stmt.Exec(new_tel, new_email, username)
	if err != nil {
		fmt.Println("cannot execute")
		return false
	}

	db.Close()
	return true
}

func DB_change_password(username, new_password string) bool {
	db, _ := dbInit()

	stmt, err := db.Prepare("UPDATE userinfo SET password=? WHERE username=?")
	if err != nil {
		fmt.Println("cannot prepare")
		return false
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(new_password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("cannot generate hashed password")
		return false
	}
	hashed_password := string(hash)

	_, err = stmt.Exec(hashed_password, username)
	if err != nil {
		fmt.Println("cannot execute")
		return false
	}
	return true
}