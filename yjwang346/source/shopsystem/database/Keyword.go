package database

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//将查询关键字存入表中
func Storekeyword(c *gin.Context,keyword string){
	stmt, err := db.Prepare("INSERT shopkeyword SET keyword=?")
	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"success":false,
			"error": "添加热门关键词失败",
			"data":err,//失败的时候留空
		})
		return
	}
	_, err = stmt.Exec(keyword)
	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"success":false,
			"error": "热门关键词添加失败",
			"data":"",//失败的时候留空
		})
		return
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"success":true,
		"error":"",
		"data":"添加热门关键词ok",	//不用传回用户的，而且没有id这一说
	})
	return
}

func GetHotkeyword(c *gin.Context) ([]string,int) {
	//返回表中keyword重复数排名前十的数值，
	//通过添加DISTINCT返回唯一记录
	rows,err := db.Query("SELECT DISTINCT * AS count FROM shopkeyword GROUP BY keyword ORDER BY count DESC LIMIT ?,?")
	var id int
	var keyword string
	store := make([]string,0)

	if err!=nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success":false,
			"error": "获得查询的信息出现错误",
			"data":"",//失败的时候留空
		})
		return store,0
	}
	for rows.Next() {
		err = rows.Scan(&id,&keyword)
		if err!=nil {
			return store,0
		}
		store =append(store,keyword)
	}
	return store,1
}

