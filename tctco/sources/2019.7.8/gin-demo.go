package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func main() {
	router := gin.Default()

	router.GET("/querystring", func (c *gin.Context){
		s_name := c.DefaultQuery("name", "Guest")
		s_age := c.Query("age")

		age, e := strconv.Atoi(s_age)
		if e == nil {
			judge(s_name, age, c)
		} else {
			c.String(http.StatusOK, "Wrong format!")
		}
	})
	router.Run(":8080")
}

func judge(name string, age int, c *gin.Context) {
	if age < 0 || age > 120 {
		c.String(http.StatusOK, "impossible age!")
	} else if age >= 0 && age <=50 {
		c.String(http.StatusOK, "user %s, you are very young!", name)
	} else {
		c.String(http.StatusOK, "user %s, you are very old!", name)
	}
}