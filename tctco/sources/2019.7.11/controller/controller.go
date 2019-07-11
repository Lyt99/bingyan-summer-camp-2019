package controller

import (
	"../model"
	jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"math"
	"net/http"
	"strconv"
)


const PAGESIZE int = 10


func Login(c *gin.Context) (interface{}, error) { //authCallback?
	username := c.PostForm("username")
	password := c.PostForm("password")

	result := model.DB_search_user(username)
	if user, ok := result.(*model.User); ok{
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"info": "login succeed!",
				"ID":        user.ID,
				"username":  user.Username,
				"password":  user.Password,
				"authority": user.Authority,
			})
			return user, nil
		} else {
			c.JSON(http.StatusOK, gin.H{"info": "wrong password!"})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{"info": "no such user!"})
	}
	return nil, nil
}


func AdminManage(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	c.JSON(200, gin.H{
		"info": "you are an admin!",
		"username": claims["username"],
		"authority": claims["authority"],
	})

	pageno, err := strconv.Atoi(c.Param("pageno"))
	if err != nil{
		pageno = 1
	}

	search_username := c.DefaultQuery("search_info_username", "")
	search_user_info(pageno, PAGESIZE, search_username, c)

	delete_username := c.Query("delete_user_username")
	if delete_username != "" {
		if authority, ok := claims["authority"].(int); ok{
			delete_user(delete_username, authority, c)
		}
	}
}


func InfoPage(c *gin.Context) {
	username := c.Param("username")
	if target_user, ok := model.DB_search_user(username).(model.User); ok {
		c.JSON(200, gin.H{
			"username":    username,
			"phonenumber": target_user.Phonenumber,
			"email":       target_user.Email,
		})
		c.SetCookie("phonenumber", target_user.Phonenumber,
			0, "/", "localhost", false, true,
		)// cannot understand path and domain
		c.SetCookie("email", target_user.Email,
			0, "/", "localhost", false, true,
		)
	}
}


func UpdateHandler(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	username := c.Param("username")
	if current_user_username := claims["username"]; current_user_username == username {
		new_tel := c.PostForm("new_tel")
		new_email := c.PostForm("new_email")
		if new_tel != "" && new_email != "" {
			updateinfo(username, new_tel, new_email, c)
		}
	} else {
		c.JSON(403, gin.H{
			"info": "you don't have the right to change this page",
		})
		c.Redirect(403, "/")
	}

}


func MainPage(c *gin.Context) {
	c.JSON(200, gin.H{
		"info": "hello world!",
	})
}


func delete_user(username string, authority int, c *gin.Context) {
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


func search_user_info(pageno, pagesize int, username string, c *gin.Context) {
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


func updateinfo (username, new_tel, new_email string, c *gin.Context) bool {
	old_tel, tel_err := c.Cookie("telephonenumber")
	old_email, email_err := c.Cookie("email")
	if tel_err != nil || email_err != nil{
		if user, ok := model.DB_search_user(username).(model.User); ok{
			old_email = user.Email
			old_tel = user.Phonenumber
		}
	}

	if old_tel == new_tel && old_email == new_email {
		return true
	} else if "" == new_tel {
		new_tel = old_tel
	} else if "" == new_email{
		new_email = old_email
	}

	return model.DB_update_info(username, new_tel, new_email)
}