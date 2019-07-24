package model

import (
	"fmt"
	"github.com/disintegration/imaging"
	"image"
	"image/color"
	"io"
	"mime/multipart"
	"onlineShopping/settings"
	"os"
	"path/filepath"
)

//this func is used to insert user info into database
func DBRegister(json RegisterJSON) error {
	hashedPassword, err := encryptPassword(json.Password)
	if err != nil {
		return err
	}

	stmt, err := DB.Prepare("INSERT INTO users (username, password, nickname, mobile, email) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(json.Username, hashedPassword, json.Nickname, json.Mobile, json.Email)
	if err != nil {
		return err
	}
	return nil
}

//this func is implemented in search commodity function
//it is used to store a word in the database
//it will create a new row if the keyword is not in the database
func saveHotWord(word string) error {
	if word == "" {
		return nil
	}
	stmt, err := DB.Prepare("INSERT INTO hotwords (word) VALUES (?) ON DUPLICATE KEY UPDATE num = num + 1")
	if err != nil {
		fmt.Print(err)
		return err
	}
	_, err = stmt.Exec(word)
	if err != nil {
		return err
	}
	return nil
}

//this func is used to insert new product into database
func DBReleaseProduct(json ProductJSON) error {
	stmt, err := DB.Prepare("INSERT INTO commodities (title, description, category, price, picture, pub_user) VALUES (?, ?, ?, ?, ?, ?)")
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

//this func is used to put a commodity into a user's collection
func DBCollectCommodity(userID, commodityID int) error {
	stmt, err := DB.Prepare("INSERT INTO collections (user_id, commodity_id) VALUES (?, ?)")
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

//this function is used to save pic
//on the server
func OSSavePic(file *multipart.FileHeader) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	file.Filename = hashFileName(file.Filename)
	path := filepath.Join(settings.ImageLocalPath, filepath.Base(file.Filename))
	dst, err := os.Create(path)
	if err != nil {
		return err
	}

	defer dst.Close()

	_, err = io.Copy(dst, src)
	if err != nil {
		return err
	}

	img, err := imaging.Open(path)
	image128 := imaging.Resize(img, 128, 128, imaging.Lanczos)
	dstThumbnail := imaging.New(128, 128, color.NRGBA{0, 0, 0, 0})
	dstThumbnail = imaging.Paste(dstThumbnail, image128, image.Pt(0, 0))

	err = imaging.Save(dstThumbnail, filepath.Join(settings.ImageLocalPath, "thumbnail_"+filepath.Base(file.Filename)))
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
