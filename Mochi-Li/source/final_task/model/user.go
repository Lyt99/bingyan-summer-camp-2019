package model

import "golang.org/x/crypto/bcrypt"

const PassWordCost = 12

type User struct {
	UserName string
	PasswordDigest string
	Nickname string
	Mobile   string
	Email    string
}

func (user *User) SetPassword(password string)error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return nil
}

func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password))
	return err == nil
}