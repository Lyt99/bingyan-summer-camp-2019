package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func GetDataBase() *sql.DB {
	url := "postgres://postgres:123456@localhost/gosql?sslmode=disable"
	log.Println(">>>> get database connection action start <<<<")
	db, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatal(err)
	}

	// 返回数据库连接
	return db
}
