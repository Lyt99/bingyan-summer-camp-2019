package model

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB

func DBInit() error {
	var err error
	db, err = sql.Open("mysql", "root:@/onlineshopping?charset=utf8")
	return err
}

func DBRegister(json RegisterJSON) error {
	hashedPassword, err := encryptPassword(json.Password)
	if err != nil {
		return err
	}

	stmt, err := db.Prepare("INSERT INTO users (username, password, nickname, mobile, email) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(json.Username, hashedPassword, json.Nickname, json.Mobile, json.Email)
	if err != nil {
		return err
	}
	return nil
}

func DBReleaseProduct(json ProductJSON) error {
	stmt, err := db.Prepare("INSERT INTO commodities (title, `desc`, category, price, picture, pub_user) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(json.Title, json.Desc, json.Category, json.Price, json.Picture, json.PubUser)
	if err != nil {
		fmt.Println(json.Title)
		fmt.Println(err)
		return err
	}

	return nil
}

func DBCollectCommodity(userID, commodityID int) error {
	stmt, err := db.Prepare("INSERT INTO collections (user_id, commodity_id) VALUES (?, ?)")
	if err != nil {
		fmt.Println(err)
		return err
	}
	_, err = stmt.Exec(userID, commodityID)
	if err != nil {
		fmt.Println(err)
		return err
	}
	changeCollectedNum(1, commodityID)
	return nil
}

func DBGetHotWords() ([]string, error) {
	words := make([]string, 0)
	rows, err := db.Query("SELECT word FROM hotwords ORDER BY num DESC")
	if err != nil {
		fmt.Print(err)
		return words, err
	}
	for rows.Next() {
		var word string
		rows.Scan(&word)
		words = append(words, word)
	}
	return words, nil
}

func saveHotWord(word string) error {
	if word == "" {
		return nil
	}
	stmt, err := db.Prepare("INSERT INTO hotwords (word) VALUES (?) ON DUPLICATE KEY UPDATE num = num + 1")
	if err != nil {
		fmt.Print(err, "301!!!")
		return err
	}
	_, err = stmt.Exec(word)
	if err != nil {
		return err
	}
	return nil
}

func getTotalViewCount(id int) (int, error) {
	rows, err := db.Query("SELECT view_count FROM commodities WHERE pub_user=?", id)
	if err != nil {
		return 0, err
	}

	cnt := 0
	for rows.Next() {
		var viewCnt int
		rows.Scan(&viewCnt)
		cnt += viewCnt
	}
	return cnt, nil
}

func getTotalCollectCount(userID int) (int, error) {
	rows, err := db.Query("SELECT collect_count FROM commodities WHERE pub_user=?", userID)
	var cnt int
	if err != nil {
		return cnt, err
	}
	for rows.Next() {
		var count int
		rows.Scan(&count)
		cnt += count
	}
	return cnt, err
}

func changeCollectedNum(num, commodityID int) error {
	stmt, err := db.Prepare("UPDATE commodities SET collect_count = collect_count + ? WHERE id=?")
	if err != nil {
		fmt.Println(err)
		return err
	}
	_, err = stmt.Exec(num, commodityID)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func encryptPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
