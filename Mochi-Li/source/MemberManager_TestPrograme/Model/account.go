package model

import (
	"crypto/md5"
	"encoding/hex"
	"gopkg.in/mgo.v2/bson"
)

func AddNewUser(u User) error {
	s := mongoSession.Clone()
	defer s.Close()

	c := s.DB(DBName).C(userCollectionName)

	return c.Insert(u)
}

func UpdateUser(userID string, u User) error {
	s := mongoSession.Clone()
	defer s.Close()

	c := s.DB(DBName).C(userCollectionName)
	return c.Update(bson.M{
		"User_ID": userID,
	}, u)
}

func DeleteUser(userID string) error {

	s := mongoSession.Clone()
	defer s.Close()

	c := s.DB(DBName).C(userCollectionName)

	return c.Remove(bson.M{
		"User_ID": userID,
	})
}

func UserExists(userID string, password string) (bool, bool, error) {
	s := mongoSession.Copy()
	defer s.Close()

	c := s.DB(DBName).C(userCollectionName)

	h2 := md5.New()
	h2.Write([]byte(password))
	password = hex.EncodeToString(h2.Sum(nil))

	result := User{}
	err := c.Find(bson.M{
		"User_ID":   userID,
		"password": password,
	}).One(&result)
	if err != nil {
		return false, false, err
	}
	return true, result.IsAdmin, err

}

func GetUserInfo(userID string) (u User, err error) {
	s := mongoSession.Copy()
	defer s.Close()

	c := s.DB(DBName).C(userCollectionName)

	query := bson.M{
		"User_ID": userID,
	}

	err = c.Find(query).One(&u)
	return
}

func GetAllInfo() (u []User, err error) {
	s := mongoSession.Copy()
	defer s.Close()

	c := s.DB(DBName).C(userCollectionName)

	err = c.Find(nil).All(&u)
	return
}
