package model

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func DBDeleteCommodity(id int) error {
	stmt, err := db.Prepare("DELETE FROM commodities WHERE id=?")
	if err != nil {
		fmt.Println(err)
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		fmt.Print(err)
		return err
	}
	return nil
}

func OSDeletePic(fileURL string) error {
	components := strings.Split(fileURL, "/")
	filename := components[len(components) - 1]
	err := os.Remove(filepath.Join("./upload", filename))
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func DBDeleteFromCollections(userID, commodityID int) error {
	stmt, err := db.Prepare("DELETE FROM collections WHERE user_id=? AND commodity_id=?")
	if err != nil {
		fmt.Println(err)
		return err
	}
	_, err = stmt.Exec(userID, commodityID)
	if err != nil {
		fmt.Println(err)
		return err
	}
	changeCollectedNum(-1, commodityID)
	return nil
}