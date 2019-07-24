package database

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"shopsystem/model"
)

func CreateCommodity(c *gin.Context,username,title,description,pricture string,price float64, category int){
	stmt, err := db.Prepare("INSERT shopcommodity SET pub_user=?, title=?,description=?, category=?,price=?,picture=?, view_count=?,collect_count=?")
	//id就不用添加了，数据库自己自动增加
	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"success":false,
			"error": "注册失败asdasd",
			"data":err,//失败的时候留空
		})
		return
	}
	zero:=0
	_, err = stmt.Exec( username,title,description, category,price,pricture,zero,zero)
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
		"error": username,
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
//获得自己发布的商品
func GetmeCommodity(c *gin.Context,username string){
	var mycommmodity model.Commodity
	row,err := db.Query("SELECT * FROM shopcommodity WHERE pub_user=?",username)
	if err!=nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success":false,
			"error": "获得商品信息出现错误",
			"data":"",//失败的时候留空
		})
		return
	}
	for row.Next(){
		err = row.Scan(&mycommmodity.Id,&mycommmodity.Pub_user,&mycommmodity.Title,&mycommmodity.Description,&mycommmodity.Category,&mycommmodity.Price,&mycommmodity.Picture,&mycommmodity.View_count,&mycommmodity.Collect_count)
	}
	// mycommmodity.Title == "" || mycommmodity.Description == ""
	if err != nil || mycommmodity.Title == ""{
		c.JSON(http.StatusBadRequest,gin.H{
			"success":mycommmodity,
			"error":err,
			"data":"您还没有发布商品",
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"success":true,
		"error":"",
		"data":gin.H{
			"id":mycommmodity.Id,//mycommmodity.Id,
			"titale":mycommmodity.Title,//mycommmodity.Title,
		},
	})
	return
}

//获得某个商品的详情，这个暂时没有要求，也不难。

//获得商品列表


func Commodities(c *gin.Context,findrequest model.Findcommodity)([]model.PostCommodity,error){
	//创建一个数据结构的数组来存储有相关分类的商品信息，考虑到后面还有keyword，所以多给了10倍
	commmodity_one := make([]model.PostCommodity,0)
	//这个就是最终的输出结果存储的位置
	//commmodity_two := make([]model.Commodity,findrequest.Limit*findrequest.Page+findrequest.Limit)
	var rows *sql.Rows
	var err error
	if findrequest.Category == 0 {
		rows, err = db.Query("SELECT id, title, price, category, picture FROM shopcommodity WHERE title LIKE ? OR description LIKE ?  ORDER BY id DESC LIMIT ?,?", "%"+findrequest.Keyword+"%", "%"+findrequest.Keyword+"%", findrequest.Page*findrequest.Limit, findrequest.Limit)
			if err!=nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"success":false,
					"error": err,
					"data":"进行商品查询时出现错误",//失败的时候留空
				})
				return commmodity_one,err
			}
	} else {
		rows, err = db.Query("SELECT id, title, price, category, picture FROM shopcommodity WHERE (category=? AND (title LIKE ? OR DESC LIKE ?)) ORDER BY id DESC LIMIT ?,? ", findrequest.Category, "'%"+findrequest.Keyword+"%'", "'%"+findrequest.Keyword+"%'", findrequest.Page*findrequest.Limit, findrequest.Limit)
		if err!=nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"success":false,
				"error": err,
				"data":"进行商品查询时出现错误",//失败的时候留空
			})
			return  commmodity_one,err
		}
	}
	for rows.Next() {
		var commodity model.PostCommodity
		err = rows.Scan(&commodity.Id, &commodity.Title, &commodity.Price, &commodity.Category, &commodity.Picture)
		if err != nil {
			fmt.Println(err)
			return commmodity_one, err
		}
		commmodity_one = append(commmodity_one, commodity)
	}
	return commmodity_one, nil
}


//获得单个商品信息
func Getone_Commodity(c *gin.Context,id int){
	var onecommmodity model.Commodity
	row,err := db.Query("SELECT * FROM shopcommodity WHERE id=?",id)
	if err!=nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success":false,
			"error": "获得商品信息出现错误",
			"data":"",//失败的时候留空
		})
		return
	}
	for row.Next(){
		err = row.Scan(&onecommmodity.Id,&onecommmodity.Pub_user,&onecommmodity.Title,&onecommmodity.Description,&onecommmodity.Category,&onecommmodity.Price,&onecommmodity.Picture,&onecommmodity.View_count,&onecommmodity.Collect_count)
	}
	if err != nil || onecommmodity.Title == ""{
		c.JSON(http.StatusBadRequest,gin.H{
			"success":onecommmodity,
			"error":err,
			"data":"获得商品信息失败123",
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"success":true,
		"error":"",
		"data":gin.H{
			"id":onecommmodity.Id,//mycommmodity.Id,
			"titale":onecommmodity.Title,//mycommmodity.Title,
			"desc":onecommmodity.Description,
			"price":onecommmodity.Price,
			"picture":onecommmodity.Picture,
			"view_count":onecommmodity.View_count,
			"collect_count":onecommmodity.Collect_count,
		},
	})
	return
}