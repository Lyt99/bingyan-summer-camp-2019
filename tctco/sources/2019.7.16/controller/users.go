package controller

import (
	"demo/model"
	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"math"
	"strconv"
)

const PAGESIZE int = 10

// This file contains the basic functions for normal users and administrators

func AdminManage(c *gin.Context) {
	authority := authCheck(c)
	switch authority {
	case 1:
		claims := jwt.ExtractClaims(c)
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
		authority, ok := claims["authority"].(float64)
		if delete_username != "" && ok {
			deleteUser(delete_username, int(authority), c)
		}

		search_username := c.DefaultQuery("search_info_username", "")
		searchUserInfo(pageno, PAGESIZE, search_username, c)

	case 0:
		c.JSON(401, gin.H{"info": "you are not an admin"})

	default:
		c.JSON(401, gin.H{"info": "you are nothing"})
	}
}

func UpdateHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	username := c.Param("username")
	if current_user_username := claims["username"]; current_user_username != username {
		c.JSON(401, gin.H{"info": "you don't have the right to change this page"})
		return
	}
	new_tel := c.PostForm("phonenumber")
	new_email := c.PostForm("email")

	if !format_check_email(new_email) {
		c.JSON(400, gin.H{"info": "incorrect email"})
		return
	}

	if !format_check_phonenumber(new_tel) {
		c.JSON(400, gin.H{"info": "incorrect phonenumber"})
		return
	}

	if model.DB_update_info(username, new_tel, new_email) {
		c.JSON(200, gin.H{"username": username, "phonenumber": new_tel, "email": new_email})
		return
	} else {
		c.JSON(500, gin.H{"info": "somehow failed"})
	}
}

func InfoPage(c *gin.Context) {
	username := c.Param("username")
	if target_user, ok := model.DB_search_user(username).(*model.User); ok {
		c.JSON(200, gin.H{
			"username":    username,
			"phonenumber": target_user.Phonenumber,
			"email":       target_user.Email,
		})
	} else {
		c.JSON(400, gin.H{"info": "no such user!"})
	}
}

func ChangePassword(c *gin.Context) {
	identity, _ := c.Get("username")
	if username, ok := identity.(string); ok {
		old_password := c.PostForm("old_password")
		new_password := c.PostForm("new_password")
		if old_password != "" && new_password != "" {
			changePassword(username, old_password, new_password, c)
		}
		return
	}
	c.JSON(401, gin.H{"info": "you don't have the authority to change password"})

}

func deleteUser(username string, authority int, c *gin.Context) {
	result := model.DB_search_user(username)
	if user, ok := result.(*model.User); ok {
		if user.Authority >= authority {
			c.JSON(401, gin.H{"info": "delete failed. you have no higher authority!"})
			return
		}
		if model.DB_delete_user(username) {
			c.JSON(200, gin.H{"info": "this user is swiped out of the db..."})
		} else {
			c.JSON(200, gin.H{"info": "somehow failed"})
		}
	}
	c.JSON(400, gin.H{"info": "no such user"})
}

func authCheck(c *gin.Context) int {
	claims := jwt.ExtractClaims(c)
	if authority, ok := claims["authority"].(float64); ok {
		return int(authority)
	}
	return -1
}

func searchUserInfo(pageno, pagesize int, username string, c *gin.Context) {
	users, cnt := model.DB_search_user_info(pageno, pagesize, username)
	max := int(math.Ceil(float64(cnt) / float64(pagesize)))
	if pageno > max {
		pageno = max
	} // seems useless?

	if users == nil {
		c.JSON(200, gin.H{"info": "no data"})
	}
	c.JSON(200, gin.H{"count": cnt, "pageno": pageno, "users": users})
}

func changePassword(username, old_password, new_password string, c *gin.Context) {
	if !formatCheckPassword(new_password) {
		c.JSON(400, gin.H{"info": "password too short!"})
		return
	}

	if user, ok := model.DB_search_user(username).(*model.User); ok {
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(old_password)); err != nil {
			c.JSON(400, gin.H{"info": "wrong password"})
			return
		}
		if model.DB_change_password(username, new_password) {
			c.JSON(200, gin.H{"info": "changed!"})
		} else {
			c.JSON(200, gin.H{"info": "somehow failed"})
		}
	}
}
