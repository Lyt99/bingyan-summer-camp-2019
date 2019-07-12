package controller

import (
	"../model"
	"fmt"
	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"math"
	"strconv"
)

const PAGESIZE int = 10

// This file contains the basic functions for normal users and administrators


func AdminManage(c *gin.Context) {
	if authority := authCheck(c); authority > 0 {
		claims := jwt.ExtractClaims(c)
		fmt.Println(claims)
		c.JSON(200, gin.H{
			"info":      "you are an admin!",
			"username":  claims["username"],
			"authority": claims["authority"],
		})

		pageno, err := strconv.Atoi(c.Param("pageno"))
		if err != nil {
			pageno = 1
		}
		delete_username := c.Query("delete_user_username")
		if delete_username != "" {
			if authority, ok := claims["authority"].(float64); ok {
				deleteUser(delete_username, int(authority), c)
			}
		}

		search_username := c.DefaultQuery("search_info_username", "")
		searchUserInfo(pageno, PAGESIZE, search_username, c)

	} else if authority == 0 {
		c.JSON(401, gin.H{
			"info": "you are not an admin",
		})
	} else {
		c.JSON(401, gin.H{
			"info": "you are nothing",
		})
	}
}


func UpdateHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	username := c.Param("username")
	if current_user_username := claims["username"]; current_user_username == username {
		old_tel, tel_err := c.Cookie("telephonenumber") // this is only a practice for cookies
		old_email, email_err := c.Cookie("email")		// it's like kind of autofill if one of the blanks is empty
		if tel_err != nil || email_err != nil{					// normally, the backend will simply refuse the request, which is safer
			if user, ok := model.DB_search_user(username).(model.User); ok{
				old_email = user.Email
				old_tel = user.Phonenumber
			}
		}
		new_tel := c.DefaultPostForm("phonenumber", old_tel)
		new_email := c.DefaultPostForm("email", old_email)


		if !format_check_email(new_email) {
			c.JSON(400, gin.H{
				"info": "incorrect email",
			})
			return;
		}

		if !format_check_phonenumber(new_tel) {
			c.JSON(400, gin.H{
				"info": "incorrect phonenumber",
			})
			return;
		}


		if new_tel == old_tel && new_email == old_email {
			c.JSON(200, gin.H{
				"username": username,
				"phonenumber": new_tel,
				"email": new_email,
			})
		} else {
			result := model.DB_update_info(username, new_tel, new_email)
			if result {
				c.JSON(200, gin.H{
					"username": username,
					"phonenumber": new_tel,
					"email": new_email,
				})
			} else {
				c.JSON(200, gin.H{
					"info": "somehow failed",
				})
			}
		}
	} else {
		c.JSON(403, gin.H{
			"info": "you don't have the right to change this page",
		})
	}

}


func InfoPage(c *gin.Context) {
	username := c.Param("username")
	if target_user, ok := model.DB_search_user(username).(*model.User); ok {
		c.SetCookie("phonenumber", target_user.Phonenumber, 0, "/", "localhost", false, true, )// cannot understand path and domain
		c.SetCookie("email", target_user.Email, 0, "/", "localhost", false, true, )
		c.JSON(200, gin.H{
			"username":    username,
			"phonenumber": target_user.Phonenumber,
			"email":       target_user.Email,
		})
	} else {
		c.JSON(404, gin.H{
			"info": "no such user!",
		})
	}
}


func ChangePassword(c *gin.Context) {
	identity, _:= c.Get("username")
	fmt.Println(identity)
	if username, ok := identity.(string); !ok {
		c.JSON(401, gin.H{
			"info": "you don't have the authority to change password",
		})
	} else {
		old_password := c.PostForm("old_password")
		new_password := c.PostForm("new_password")
		if old_password != "" && new_password != "" {
			changePassword(username, old_password, new_password, c)
		}
	}
}



func deleteUser(username string, authority int, c *gin.Context) {
	result := model.DB_search_user(username)
	if user, ok := result.(*model.User); ok {
		if user.Authority >= authority {
			c.JSON(200, gin.H{
				"info": "delete failed. you have no higher authority!",
			})
		} else {
			model.DB_delete_user(username)
			c.JSON(200, gin.H{
				"info": "this user is swiped out of the db...",
			})
		}
	} else {
		c.JSON(200, gin.H{
			"info": "no such user",
		})
	}
}


func authCheck(c *gin.Context) int {
	claims := jwt.ExtractClaims(c)
	if authority, ok := claims["authority"].(float64); ok{
		return int(authority)
	}
	return -1
}


func searchUserInfo(pageno, pagesize int, username string, c *gin.Context) {
	users, cnt := model.DB_search_user_info(pageno, pagesize, username)
	max := int(math.Ceil(float64(cnt) / float64(pagesize)))
	if pageno > max{
		pageno = max
	}// seems useless?

	if users != nil {
		c.JSON(200, gin.H{
			"count": cnt,
			"pageno": pageno,
			"users": users,
		})
	} else {
		c.JSON(200, gin.H{
			"info": "no data",
		})
	}
}


func changePassword(username, old_password, new_password string, c *gin.Context) {
	if !formatCheckPassword(new_password){
		c.JSON(400, gin.H{
			"info": "password too short!",
		})
		return
	}
	if user, ok := model.DB_search_user(username).(*model.User); ok {
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(old_password)); err == nil{
			model.DB_change_password(username, new_password)
			c.JSON(200, gin.H{
				"info": "changed!",
			})
		} else {
			c.JSON(400, gin.H{
				"info": "wrong password",
			})
		}
	}
}