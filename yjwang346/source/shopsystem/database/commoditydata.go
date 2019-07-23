package database

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func CreateCommodity(c *gin.Context,pub_user,title,description,pricture string,price float64, category int){
	stmt, err := db.Prepare("INSERT shopcommodity SET pub_user=?, title=?,description=?, price=?,category=?,pricture=?, view_count=?,collect_count=?")
	//id就不用添加了，数据库自己自动增加
	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"success":false,
			"error": "注册失败asdasd",
			"data":"",//失败的时候留空
		})
		return
	}
	zero:=0
	_, err = stmt.Exec(pub_user,title,description, price, category,pricture,zero,zero)
	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"success":false,
			"error": "添加商品失败",
			"data":"",//失败的时候留空
		})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"success":true,
		"error": "",
		"data":"添加商品ok",	//不用传回用户的，而且没有id这一说
	})
	return
}

//删除相应的商品
func DeleteCommo(c *gin.Context,id int){

	_, err := db.Prepare("DELETE FROM shopcommodity WHERE id=?")
	if err != nil {
		fmt.Println("1")
		log.Fatalln(err)
	}
	return
}