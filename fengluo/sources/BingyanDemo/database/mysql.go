package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)
var db *sql.DB

func InitMysql() {
	fmt.Println("InitMysql....")
	if db == nil {
		var err error
		db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/user_info")
		if err != nil {
			log.Fatal(err);
			return
		}
		CreateTableWithUser()
	}
}
func CreateTableWithUser() {
	sql := `CREATE TABLE IF NOT EXISTS information(
        id INT(20) PRIMARY KEY AUTO_INCREMENT NOT NULL,
		user_id VARCHAR(40),
		user_password VARCHAR(128),
		user_name VARCHAR(20),
        user_phone VARCHAR(40),
		user_email VARCHAR(40),
        user_admin TINYINT(1)
        );`

	ModifyDB(sql)
}

func FindDB(sql string) *sql.Row{
	return db.QueryRow(sql)
}

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