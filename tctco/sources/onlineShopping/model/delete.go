package model

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func DBDeleteCommodity(id int) error {
	stmt, err := DB.Prepare("DELETE FROM commodities WHERE id=?")
	if err != nil {
		fmt.Println(err)
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		fmt.Print(err)
		return err
	}
	userIDs, _ := getCommodityCollectors(id)
	for _, v := range userIDs {
		strID := strconv.Itoa(v)
		addInfo("商品"+strID+"已被删除！", v)
	}
	return nil
}

//this func is used to delete the pic related to the commodity
func OSDeletePic(fileURL string) error {
	components := strings.Split(fileURL, "/")
	filename := components[len(components)-1]
	err := os.Remove(filepath.Join("./upload", filename))
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

//this func is used to delete a commodity from a collection
func DBDeleteFromCollections(userID, commodityID int) error {
	stmt, err := DB.Prepare("DELETE FROM collections WHERE user_id=? AND commodity_id=?")
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
