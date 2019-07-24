package model

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	"onlineShopping/settings"
	"path/filepath"
	"strconv"
	"time"
)

var DB *sql.DB

func DBInit() error {
	var err error
	DB, err = sql.Open("mysql", settings.DataBasePath)
	return err
}

//this func is used to get the view count of a commodity
func getTotalViewCount(id int) (int, error) {
	rows, err := DB.Query("SELECT view_count FROM commodities WHERE pub_user=?", id)
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

//this func is used to get the total collect number of a user
func getTotalCollectCount(userID int) (int, error) {
	rows, err := DB.Query("SELECT collect_count FROM commodities WHERE pub_user=?", userID)
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

//this func is used to change the collect number of a product
//the first parameter can be +1 or -1, which means collect
//or delete from the collection
func changeCollectedNum(num, commodityID int) error {
	stmt, err := DB.Prepare("UPDATE commodities SET collect_count = collect_count + ? WHERE id=?")
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

//this is used to hash the password
func encryptPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

//this func also adds a time stamp to distinguish
//pics with the same name
func hashFileName(filename string) string {
	m := md5.New()
	filename = filename
	m.Write([]byte(filename))
	extension := filepath.Ext(filename)

	return hex.EncodeToString(m.Sum(nil)) + strconv.FormatInt(time.Now().Unix(), 10) + extension
}

func addInfo(info string, receiverID int) error {
	stmt, err := DB.Prepare("INSERT INTO info (info, receiver_id) VALUES (?, ?)")
	if err != nil {
		fmt.Println(err)
		return err
	}

	_, err = stmt.Exec(info, receiverID)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
