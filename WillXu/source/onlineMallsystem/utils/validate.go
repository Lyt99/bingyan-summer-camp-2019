package utils

import (
	"regexp"
)

const mobileRe  = `^1([38][0-9]|14[57]|5[^4])\d{8}$`
const emailRe  = "^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+(.[a-zA-Z0-9_-])+"

func CheckMobile(mobileNum string) bool {
	reg := regexp.MustCompile(mobileRe)
	return reg.MatchString(mobileNum)
}

func CheckEmail(emailAdd string) bool {
	reg := regexp.MustCompile(emailRe)
	return reg.MatchString(emailAdd)
}