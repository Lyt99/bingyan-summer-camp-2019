package model

import (
	"database/sql"
	"fmt"
	"strings"
)

//this func search a user by username
//can be used in login part
func DBSearchUser(username string) (User, error) {
	var user User

	row := DB.QueryRow("SELECT * FROM users WHERE username=?", username)

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

	row := DB.QueryRow("SELECT * FROM users WHERE id=?", id)

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

//this func search commodities by keywords
//and categories. The search part is realized by
//mysql commands
func DBSearchCommodities(json SearchJSON) ([]CommodityBrief, error) {
	commodities := make([]CommodityBrief, 0)
	var rows *sql.Rows
	var err error

	if json.Category == 0 {
		rows, err = DB.Query("SELECT id, title, price, category, picture FROM commodities WHERE title LIKE ? OR description LIKE ?  ORDER BY id DESC LIMIT ?,?", "%"+json.Keyword+"%", "%"+json.Keyword+"%", json.PageNo*json.PageSize, json.PageSize)
		if err != nil {
			fmt.Println(err)
			return commodities, err
		}
	} else {
		rows, err = DB.Query("SELECT id, title, price, category, picture FROM commodities WHERE (category=? AND (title LIKE ? OR description LIKE ?)) ORDER BY id DESC LIMIT ?,? ", json.Category, "'%"+json.Keyword+"%'", "'%"+json.Keyword+"%'", json.PageNo*json.PageSize, json.PageSize)
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
		components := strings.Split(commodity.Picture, "/")
		components[len(components)-1] = "thumbnail_" + components[len(components)-1]
		commodity.Thumbnail = strings.Join(components, "/")
		commodities = append(commodities, commodity)
	}
	saveHotWord(json.Keyword)
	return commodities, nil
}

//this func is used to search commodities
//that belongs to a particular user
func DBSearchUserCommodities(userID int) ([]MyCommodity, error) {
	myCommodities := make([]MyCommodity, 0)

	rows, err := DB.Query("SELECT id, title FROM commodities WHERE pub_user=?", userID)
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

//this func is used to search a single
//commodity by its id. This can be used to look
//for detailed commodity information
func DBSearchCommodity(id int) (Commodity, bool, error) {
	var commodity Commodity
	rows, err := DB.Query("SELECT title, description, category, price, picture, pub_user, view_count, collect_count FROM commodities WHERE id=?", id)
	if err != nil {
		fmt.Println(err)
		return commodity, false, err
	}
	for rows.Next() {
		rows.Scan(&commodity.Title, &commodity.Desc, &commodity.Category, &commodity.Price,
			&commodity.Picture, &commodity.PubUser, &commodity.ViewCount, &commodity.CollectCount)
		stmt, err := DB.Prepare("UPDATE commodities SET view_count = view_count + 1 WHERE id=?")
		if err != nil {
			fmt.Println(err)
			return commodity, true, err
		}
		stmt.Exec(id)
		return commodity, true, nil
	}
	return commodity, false, nil
}

//this func is used to search all collected
//commodities. The search part is realized by
//INNER JOIN 2 tables in mysql
func DBSearchCollections(userID int) ([]MyCollection, error) {
	collection := make([]MyCollection, 0)
	rows, err := DB.Query("SELECT cm.id, cm.title FROM collections cl INNER JOIN commodities cm ON cl.commodity_id=cm.id WHERE cl.user_id = ?", userID)
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

//this func is used to get all hot words from the database
func DBGetHotWords() ([]string, error) {
	words := make([]string, 0)
	rows, err := DB.Query("SELECT word FROM hotwords ORDER BY num DESC")
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

func DBGetUserMessages(userID int) ([]Message, error) {
	messages := make([]Message, 0)
	rows, err := DB.Query("SELECT id, sender_id, info FROM info WHERE receiver_id = ? AND is_read = 0", userID)
	if err != nil {
		fmt.Println(err)
		return messages, err
	}

	stmt, err := DB.Prepare("UPDATE info SET is_read = 1 WHERE id = ? AND is_read = 0")

	for rows.Next() {
		var id int
		var message Message
		rows.Scan(&id, &message.SenderID, &message.Info)
		stmt.Exec(id)
		messages = append(messages, message)
	}
	return messages, nil
}

func getCommodityCollectors(commodityID int) ([]int, error) {
	userIDs := make([]int, 0)
	rows, err := DB.Query("SELECT user_id FROM collections WHERE commodity_id=?", commodityID)
	if err != nil {
		fmt.Println(err)
		return userIDs, err
	}

	for rows.Next() {
		var userID int
		rows.Scan(&userID)
		userIDs = append(userIDs, userID)
	}
	return userIDs, nil
}
