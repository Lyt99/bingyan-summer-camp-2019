package model

import (
	"database/sql"
	"fmt"
)

func DBSearchUser(username string) (User, error) {
	var user User

	row := db.QueryRow("SELECT * FROM users WHERE username=?", username)

	err := row.Scan(&user.ID, &user.Username, &user.Nickname, &user.Password, &user.Mobile, &user.Email)
	if err != nil {
		fmt.Println(err)
		return user, err
	}

	user.TotalViewCount, err = getTotalViewCount(user.ID)
	if err != nil {
		fmt.Println(err)
		return user, err
	}

	user.TotalCollectCount, err = getTotalCollectCount(user.ID)
	if err != nil {
		fmt.Println(err)
		return user, err
	}
	return user, nil
}

func DBSearchUserByID(id int) (User, error) {
	var user User

	row := db.QueryRow("SELECT * FROM users WHERE id=?", id)

	err := row.Scan(&user.ID, &user.Username, &user.Nickname, &user.Password, &user.Mobile, &user.Email)
	if err != nil {
		fmt.Println(err)
		return user, err
	}

	user.TotalViewCount, err = getTotalViewCount(user.ID)
	if err != nil {
		fmt.Println(err)
		return user, err
	}

	user.TotalCollectCount, err = getTotalCollectCount(user.ID)
	if err != nil {
		fmt.Println(err)
		return user, err
	}
	return user, nil
}

func DBSearchCommodities(json SearchJSON) ([]CommodityBrief, error) {
	commodities := make([]CommodityBrief, 0)
	var rows *sql.Rows
	var err error

	if json.Category == 0 {
		rows, err = db.Query("SELECT id, title, price, category, picture FROM commodities WHERE title LIKE ? OR `desc` LIKE ?  ORDER BY id DESC LIMIT ?,?", "%"+json.Keyword+"%", "%"+json.Keyword+"%", json.PageNo*json.PageSize, json.PageSize)
		if err != nil {
			fmt.Println(err)
			return commodities, err
		}
	} else {
		rows, err = db.Query("SELECT id, title, price, category, picture FROM commodities WHERE (category=? AND (title LIKE ? OR DESC LIKE ?)) ORDER BY id DESC LIMIT ?,? ", json.Category, "'%"+json.Keyword+"%'", "'%"+json.Keyword+"%'", json.PageNo*json.PageSize, json.PageSize)
		if err != nil {
			fmt.Println(err)
			return commodities, err
		}
	}
	for rows.Next() {
		var commodity CommodityBrief
		err = rows.Scan(&commodity.ID, &commodity.Title, &commodity.Price, &commodity.Category, &commodity.Picture)
		if err != nil {
			fmt.Println(err)
			return commodities, err
		}
		commodities = append(commodities, commodity)
	}
	saveHotWord(json.Keyword)
	return commodities, nil
}

func DBSearchUserCommodities(userID int) ([]MyCommodity, error) {
	myCommodities := make([]MyCommodity, 0)

	rows, err := db.Query("SELECT id, title FROM commodities WHERE pub_user=?", userID)
	if err != nil {
		fmt.Println(err)
		return myCommodities, err
	}
	for rows.Next() {
		var commodity MyCommodity
		rows.Scan(&commodity.ID, &commodity.Title)
		myCommodities = append(myCommodities, commodity)
	}
	return myCommodities, nil
}

func DBSearchCommodity(id int) (Commodity, bool, error) {
	var commodity Commodity
	rows, err := db.Query("SELECT title, `desc`, category, price, picture, pub_user, view_count, collect_count FROM commodities WHERE id=?", id)
	if err != nil {
		fmt.Println(err)
		return commodity, false, err
	}
	for rows.Next() {
		rows.Scan(&commodity.Title, &commodity.Desc, &commodity.Category, &commodity.Price,
			&commodity.Picture, &commodity.PubUser, &commodity.ViewCount, &commodity.CollectCount)
		stmt, err := db.Prepare("UPDATE commodities SET view_count = view_count + 1 WHERE id=?")
		if err != nil {
			fmt.Println(err)
			return commodity, true, err
		}
		stmt.Exec(id)
		return commodity, true, nil
	}
	return commodity, false, nil
}

func DBSearchCollections(userID int) ([]MyCollection, error) {
	collection := make([]MyCollection, 0)
	rows, err := db.Query("SELECT cm.id, cm.title FROM collections cl INNER JOIN commodities cm ON cl.commodity_id=cm.id WHERE cl.user_id = ?", userID)
	if err != nil {
		fmt.Println(err)
		return collection, err
	}
	for rows.Next() {
		var fav MyCollection
		rows.Scan(&fav.ID, &fav.Title)
		collection = append(collection, fav)
	}
	return collection, nil
}