package main
import (
"database/sql"
	"fmt"
	_ "github.com/Go-SQL-Driver/MySQL"
"github.com/gin-gonic/gin"
"log"
"net/http"
)
var db *sql.DB
var err0 error
type Login struct {
Username    string `form:"username  " json:"username  " binding:"required"`
Password  string `form:"password" json:"password" binding:"required"`
Nickname  string   `form:"nickname" json:"nickname" `
Mobile   string    `form:"mobile" json:"mobile"`
Email string   `form:"email" json:"email" `
}
type Sell struct {
	Page     int    `form:"page"json:"page"    `
	Limit    int    `form:"limit"json:"limit"binding:"required"`
	Category int    `form:"category"json:"category"`
	Keyword  string `form:"keyword"json:"keyword"`
}
type Goods struct{
     Desc   string  `json:"desc"`
     Price int `json:"price"`
     Picture int `json:"picture"`
     PubUser string  `json:"pubuser"`
     Title string `json:"title"`
	Category int    `form:"category"json:"category"`
}
func main(){
r := gin.Default()
db, err0 = sql.Open("mysql","root:1751930896@tcp(127.0.0.1:3306)/test2")
if err0 != nil {
log.Fatalln(err0)
}
r.POST("/user", func(c *gin.Context)  {
var json Login
	json.Nickname=c.PostForm("nickname")
	json.Username=c.Query("username")
	json.Password =c.Query("password")
	json.Mobile = c.PostForm("mobile")
	json.Email= c.PostForm("email")
	fmt.Println(json)
	if json.Username != "" && json.Password != "" {
		json.Password=encrypt(json.Password)
err2,_:=db.Exec("INSERT INTO test2.user (username,password,email,nickname,mobile) VALUES(?,?,?,?,?)", json.Username, json.Password, json.Email, json.Nickname, json.Mobile)
if err2!=nil{
c.JSON(http.StatusOK, gin.H{
	"success": true,
	"error": "",
	"data": "ok"}) }else{
c.JSON(http.StatusNotFound, gin.H{ "success": false,
	"error": "用户名已存在",
	"data": "null"})
}
} else {
c.JSON(http.StatusNotFound, gin.H{ "success": false,
	"error": "请正确输入信息",
	"data": "null"})
}
})
r.POST("/login", func(c *gin.Context) {
		var json Login
		if err := c.ShouldBindJSON(&json); err == nil {
			err1:= db.QueryRow("SELECT password FROM test2.user WHERE username=?;", json.Username).Scan(&json.Password)
			if err1==nil{
				c.JSON(http.StatusOK, gin.H{"status": "You are login","message":json})
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"status": "Unauthorized"})
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})
r.GET("/commodities",func(c *gin.Context) {
       var goods Sell
	   if err := c.ShouldBindJSON(&goods); err == nil {
		   for i := goods.Page* goods.Limit; i <=( goods.Page+1)*goods.Limit; i++ {
		   	if goods.Category!=0 {
				rows, err5 := db.Query("SELECT id,title,price,cano,picture FROM test2.sells where id =? and cano=?;", i, goods.Category)
			   checkErr(err5)
			   for rows.Next() {
				   var id int
				   var title string
				   var price int
				   var cano int
				   var picture string
				   err5 = rows.Scan(&id, &title, &price, &cano, &picture)
				   checkErr(err5)
				   c.JSON(http.StatusOK, gin.H{"id ": id,
					   "title ":         title,
					   "price ":     price,
					   "category":         cano,
					   "picture ": picture})
			   }}else {
				rows, err5 := db.Query("SELECT id,title,price,cano,picture FROM test2.sells where id =? ", i)
				checkErr(err5)
				for rows.Next() {
					var id int
					var title string
					var price int
					var cano int
					var picture string
					err5 = rows.Scan(&id, &title, &price, &cano, &picture)
					checkErr(err5)
					c.JSON(http.StatusOK, gin.H{"id ": id,
						"title ":         title,
						"price ":     price,
						"category":         cano,
						"picture ": picture})}
		   }}
	   }else{
	   	c.JSON(http.StatusNotFound,"参数不合法")
	   }
})
r.GET("/commodities/hot",func(c *gin.Context) {
})
r.POST("/commodities",func(c *gin.Context){
  var goods Goods
	if err := c.ShouldBindJSON(&goods); err == nil {
		err2,_:=db.Exec("INSERT INTO test2.sells (pub_user,title,price,cano,picture,`desc`) VALUES(?,?,?,?,?,?)", goods.PubUser, goods.Title,goods.Price, goods.Category,goods.Picture,goods.Desc)
		if err2!=nil{
			c.JSON(http.StatusOK, gin.H{
				"success": true,
				"error": "",
				"data": "ok"}) }else{
			c.JSON(http.StatusNotFound, gin.H{ "success": false,
				"error": "error",
				"data": "null"})
		}
	} else {
		c.JSON(http.StatusNotFound, gin.H{ "success": false,
			"error": "请正确输入信息",
			"data": "null"})
	}
})
r.GET("commodity/:id",func(c *gin.Context){
	var  id int
	id=c.GetInt("id")
	rows, err5 := db.Query("SELECT pub_user,title,`desc`,cano,price,picture FROM test2.sells where id =? ", id)
	checkErr(err5)
	for rows.Next() {
		var pubuser string
		var title int
		var desc string
		var price int
		var cano int
		var picture string
		err5 = rows.Scan(&pubuser, &title, &desc, &cano, &price,&picture)
		checkErr(err5)
		c.JSON(http.StatusOK, gin.H{"pub_user ": pubuser,
			"title ":         title,
			"desc": desc,
			"category":         cano,
			"price ":     price,
			"picture ": picture})}
})
r.GET("DELETE/commodity/:id",func(c *gin.Context){
   id:=c.GetInt("id")
   username:=c.Query("username")
	rows,err3 := db.Query(
		"SELECT pub_user FROM test2.sells WHERE id=?;",id)
	checkErr(err3)
	for rows.Next() {
		var pubuser string
		err4:= rows.Scan(&pubuser)
		checkErr(err4)
 if username==pubuser{
	_,err3 = db.Query("SELECT * FROM test2.user WHERE username=?;",pubuser)
	if err3 != nil {
		err4, _ := db.Exec("delete from test2.sells where id= ?;", id)
		if err4 != nil {
			c.JSON(http.StatusOK, gin.H{
				"success": true,
				"error":   "",
				"data":    "ok"})
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"error":   "该商品不存在或者该用户不存在",
				"data":    "null"})
		}
	}else {c.JSON(http.StatusNotFound, gin.H{
		"success": false,
		"error":   "该商品并不是该用户",
		"data":    "null"})}
	}
	}
})
r.GET("GET/me", func(c *gin.Context){
	var json Login
	json.Username=c.Query("username")
		rows, err5 := db.Query("SELECT username,nickname,mobile,email FROM test2.user where username=?;",json.Username )
		checkErr(err5)

		for rows.Next() {
			var username string
			var nickname string
			var mobile string
			var email string
			err5 = rows.Scan(&username, &nickname, &nickname, &mobile, &email)
			checkErr(err5)
			c.JSON(http.StatusOK, gin.H{"username ": username,
				"nickname ":         nickname,
				"mobile":          mobile,
				"email: ": email})
		}
})
r.GET("POST/me",func(c*gin.Context){
	var json Login
	json.Username=c.Query("username")
	json.Nickname=c.Query("nickname")
	json.Password = encrypt(c.Query("Password"))
	json.Mobile = c.Query("mobile")
	json.Email= c.Query("email")
	//if err5,_:=db.Exec("select * FROM test.USER where ID=?;",json.ID);err5==nil {
	err5,_:=db.Exec("INSERT INTO test2.user (username,password,email,nickname,mobile) VALUES(?,?,?,?,?,?)", json.Username, json.Password, json.Email, json.Nickname, json.Mobile)
	if err5==nil{
		if json.Password != "" {
			if err1, _ := db.Exec("update test2.user set password = ? where username= ?;", json.Password, json.Username); err1 != nil {
				c.String(http.StatusOK, " Update password Success")
			} else {
				c.String(http.StatusNotFound, " Password Failed")
			}
		}
		if json.Email != "" {
			if err2, _ := db.Exec("update test2.user set email= ? where username= ?;", json.Email, json.Username); err2 != nil {
				c.String(http.StatusOK, " Update EmailAddress Success")
			} else {
				c.String(http.StatusNotFound, " Email Failed")
			}
		}
		if json.Mobile != "" {
			if err3, _ := db.Exec("update test2.user set mobile= ? where username = ?;", json.Mobile, json.Username); err3 != nil {
				c.String(http.StatusOK, " Update Tel Success")
			} else {
				c.String(http.StatusNotFound, " Mobile Failed")
			}
		}
		if json.Nickname != "" {
			if err4, _ := db.Exec("update test2.user set nickname = ? where username = ?;", json.Nickname, json.Username); err4 != nil {
				c.String(http.StatusOK, " Update Name Success")
			} else {
				c.String(http.StatusNotFound, " Nickname Failed")
			}
		}
	}else {
		db.Exec("Delete from test2.USER where username=?;",json.Username)
		c.String(http.StatusNotFound,"ID not exist")
	}
})
r.GET("GET/me/commodities",func(c*gin.Context){
	pubuser:=c.Query("username")
	rows,err1:=db.Query("select id,title from test2.sells where pub_user=?;",pubuser)
	checkErr(err1)
	for rows.Next (){
    var id int
    var title string
    err1=rows.Scan(&id,&title)
    checkErr(err1)
    c.JSON(http.StatusOK,gin.H{
		"success": true,
		"error": "",
		"id": id,
		"title": title})
	}
})
type Se struct {
	ID int`json:"id"`
	Username string  `json:"username"`
}
r.GET("get/me/collections",func(c*gin.Context){
      var json Se
      json.Username=c.Query("username")
      rows,err1:=db.Query("select ID from test2.collection WHERE user=?;",json.Username)
      checkErr(err1)
      for rows.Next(){
      	var ID int
      	err1=rows.Scan(&ID)
        checkErr(err1)
      	rows,err1=db.Query("select title from test2.sells where id=?;",ID)
      	checkErr(err1)
      	for rows.Next(){
			var title string
			err1=rows.Scan(&title)
			c.JSON(http.StatusOK,gin.H{
				"success": true,
				"error": "",
				"id": ID,
				"title": title})
		}
	  }
})
r.POST("POST/me/collections",func(c*gin.Context){
	var json Se
	json.Username=c.Query("username")
	json.ID=c.GetInt("id")
	err1,_:=db.Exec("INSERT INTO test2.collection (ID,user) VALUES(?,?)", json.ID, json.Username)
	if err1!=nil{
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"error": "",
			"data": "ok"}) }else{
		c.String(http.StatusNotFound, "商品不存在或用户未登陆")}
})
r.GET("DELETE/me/collections",func(c*gin.Context){
	id :=c.GetInt("id")
	user:=c.Query("username")
	err3 := db.QueryRow(
		"SELECT * FROM test2.collection WHERE ID=? and user=?;",id,user)
	if err3 != nil {
		err4,_:=db.Exec("delete from test2.collection where ID = ? and user=?;",id,user)
		if err4!=nil{
			c.JSON(http.StatusOK, gin.H{
				"success": true,
				"error": "",
				"data": "ok"})} else {
			c.String(http.StatusNotFound, " Failed")}
	}else {
		c.String(http.StatusNotFound, " 商品不存在收藏中")
	}
})
r.GET ("GET?/user/:id",func(c*gin.Context){
	var json Login
	json.Username=c.Query("username")
	rows, err5 := db.Query("SELECT username,nickname,mobile,email FROM test2.user where username=?;",json.Username )
	checkErr(err5)

	for rows.Next() {
		var username string
		var nickname string
		var mobile string
		var email string
		err5 = rows.Scan(&username, &nickname, &nickname, &mobile, &email)
		checkErr(err5)
		c.JSON(http.StatusOK, gin.H{"username ": username,
			"nickname ":         nickname,
			"mobile":          mobile,
			"email: ": email})
	}
})
r.Run(":8080")
}
func StringToRuneArr(s string, arr []rune) {
	src := []rune(s)
	for i, v := range src {
		if i >= len(arr) {
			break
		}
		arr[i] = v+32
	}
}

func encrypt(str string ) string {
	var arr [10]rune
	StringToRuneArr(str, arr[:])
	st := string(arr[:])
	return st
}
func checkErr(err error) {
	if err != nil {
		fmt.Println("ID not exist")
		panic(err)
	}
}

