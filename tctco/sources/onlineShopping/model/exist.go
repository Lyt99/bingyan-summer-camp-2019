package model

import "fmt"

func DBIsCollectionExist(userID, commodityID int) (bool, error) {
	rows, err := DB.Query("SELECT * FROM collections WHERE user_id=? AND commodity_id=?", userID, commodityID)
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	return rows.Next(), nil
}

func DBIsCommodityExist(id int) (int, error) {
	rows, err := DB.Query("SELECT pub_user FROM commodities WHERE id=?", id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	for rows.Next() {
		var pubID int
		rows.Scan(&pubID)
		return pubID, nil
	}
	return 0, nil
}

func DBIsUserExist(username string) (bool, error) {
	rows, err := DB.Query("SELECT * FROM users WHERE username=?", username)
	if err != nil {
		fmt.Println("DB query fail")
		return false, err
	}

	return rows.Next(), err
}
