
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)
type Login struct {
	ID     string `form:"ID" json:"ID" binding:"required"`
	Password  string `form:"Password" json:"Password" binding:"required"`
	Name  string   `form:"Name" json:"Name" `
	Tel   string    `form:"Tel" json:"Tel"`
	EmailAddress string   `form:"EmailAddress" json:"EmailAddress" `
	IsAdmin string   `form:"IsAdmin" json:"IsAdmin" `
	DeID     string   `form:"DeID" json:"DeID" `
	SeID     string   `form:"SeID" json:"SeID" `
}
var db *sql.DB
//func encrypt(arr string)  string {
//	//var a[] string
//	//strings.Split(arr,"" )
//	// for i:=0;;i++ {
//		//a[i]=a[i]+'1'
//	//}
//	arr =arr+"ahasdfgjjkl"
//	return  arr
//}
func StringToRuneArr(s string, arr []rune) {
	src := []rune(s)
	for i, v := range src {
		if i >= len(arr) {
			break
		}
		arr[i] = v+32
	}
}
func StringToRuneArr1(s string, arr []rune) {
	src := []rune(s)
	for i, v := range src {
		if i >= len(arr) {
			break
		}
		arr[i] = v-32
	}
}

func encrypt(str string ) string{
	//str := "ABCDEF"
	var arr [10]rune
	StringToRuneArr(str, arr[:])
	//fmt.Println(string(arr[:]))
	st:= string(arr[:])
	return st
}
func decrypt(str string)  string {
	var arr [10]rune
	StringToRuneArr1(str, arr[:])
	//fmt.Println(string(arr[:]))
	st:= string(arr[:])
	return st
}
var err0 error
func main() {
	r := gin.Default()
	db, err0 = sql.Open("mysql","root:1751930896@tcp(127.0.0.1:3306)/test")
	if err0 != nil {
		log.Fatalln(err0)
	}
	defer db.Close()
	//CreateTable(db)
	//stmt, err := db.Prepare("CREATE TABLE USER ( ID  char(10) UNIQUE,password char(20) NOT NULL,name char(20),tel char(15),e_mail_address char(20),IsAdmin char(1) check ( IsAdmin in ('T', 'F') ));")
	//登陆
	r.GET("/login", func(c *gin.Context) {
		var json Login
		if err := c.ShouldBindJSON(&json); err == nil {
			err1:= db.QueryRow("SELECT password FROM test.USER WHERE ID=?;", json.ID).Scan(&json.Password)
			if err1==nil{
				//c.String(http.StatusOK,"you are login\n",json)
				c.JSON(http.StatusOK, gin.H{"status": "You are login","message":json})
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"status": "Unauthorized"})
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})
//注册
	r.POST("/welcome", func(c *gin.Context)  {
		  var json Login
		  json.Name=c.PostForm("Name")
		  json.ID=c.Query("ID")
			json.Password = encrypt(c.Query("Password"))
			json.Tel = c.PostForm("Tel")
			json.EmailAddress = c.PostForm("EmailAddress")
			json.IsAdmin = c.PostForm("IsAdmin")
			fmt.Println(json)
			if json.ID != "" && json.Password != "" {
			   err2,_:=db.Exec("INSERT INTO test.USER (ID,password,e_mail_address,name,tel,IsAdmin) VALUES(?,?,?,?,?,?)", json.ID, json.Password, json.EmailAddress, json.Name, json.Tel, json.IsAdmin)
			   if err2!=nil{
				c.String(http.StatusOK, " Success") }else{
				   c.String(http.StatusNotFound, "Failed")
			   }
			} else {
				c.String(http.StatusNotFound, "Failed")
			}
		  // }
	})
//
	r.GET("login/Delete",func(c *gin.Context) {
		var json Login
		json.DeID = c.Query("DeID")
		json.IsAdmin = c.Query("IsAdmin")
		if json.IsAdmin == "T" {
			err3 := db.QueryRow(
			"SELECT * FROM test.USER WHERE ID=?;",json.DeID)
			if err3 != nil {
				err4,_:=db.Exec("delete from test.USER where ID = ?;",json.DeID)
				if err4!=nil{
					c.String(http.StatusOK, " Success")} else {
					c.String(http.StatusNotFound, " Failed")}
			}
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "Unauthorized"})
		}
	})
	r.GET("login/Update",func(c *gin.Context) {
		var json Login
		json.Name=c.Query("Name")
		json.ID=c.Query("ID")
		json.Password = encrypt(c.Query("Password"))
		json.Tel = c.Query("Tel")
		json.EmailAddress = c.Query("EmailAddress")
		json.IsAdmin = c.Query("IsAdmin")
		//if err5,_:=db.Exec("select * FROM test.USER where ID=?;",json.ID);err5==nil {
		err5,_:=db.Exec("INSERT INTO test.USER (ID,password,e_mail_address,name,tel,IsAdmin) VALUES(?,?,?,?,?,?)", json.ID, json.Password, json.EmailAddress, json.Name, json.Tel, json.IsAdmin)
		if err5==nil{
			fmt.Println(9952)
			if json.Password != "" {
				if err1, _ := db.Exec("update test.USER set password = ? where ID = ?;", json.Password, json.ID); err1 != nil {
					c.String(http.StatusOK, " Update password Success")
				} else {
					c.String(http.StatusNotFound, " Password Failed")
				}
			}
			if json.EmailAddress != "" {
				if err2, _ := db.Exec("update test.USER set e_mail_address = ? where ID = ?;", json.EmailAddress, json.ID); err2 != nil {
					c.String(http.StatusOK, " Update EmailAddress Success")
				} else {
					c.String(http.StatusNotFound, " EmailAddress Failed")
				}
			}
			if json.Tel != "" {
				if err3, _ := db.Exec("update test.USER set tel= ? where ID = ?;", json.Tel, json.ID); err3 != nil {
					c.String(http.StatusOK, " Update Tel Success")
				} else {
					c.String(http.StatusNotFound, " TEL Failed")
				}
			}
			if json.Name != "" {
				if err4, _ := db.Exec("update test.USER set name = ? where ID = ?;", json.Name, json.ID); err4 != nil {
					c.String(http.StatusOK, " Update Name Success")
				} else {
					c.String(http.StatusNotFound, " Name Failed")
				}
			}
		}else {
			db.Exec("Delete from test.USER where ID=?;",json.ID)
			c.String(http.StatusNotFound,"ID not exist")
		}
})
	r.POST("login/Select",func(c *gin.Context) {
		var json Login
		json.ID=c.Query("ID")
		json.Password = encrypt(c.Query("Password"))
		json.IsAdmin = c.PostForm("IsAdmin")
		json.SeID = c.PostForm("SeID")
	      if json.IsAdmin == "T"{
			  rows, err5 := db.Query("SELECT * FROM test.USER where ID=?;",json.SeID )
			  checkErr(err5)

			  for rows.Next() {
				  var ID string
				  var name string
				  var password string
				  var tel string
				  var EmailAddress string
				  var IsAdmin string
				  err5 = rows.Scan(&ID,  &password,&name, &tel, &EmailAddress, &IsAdmin)
				  checkErr(err5)
				  c.JSON(http.StatusOK,gin.H{"ID is: ": ID,
				  	                         "name is: ": name,
				                             "password is: ":decrypt(password),
				                             "Tel is: ": tel,
				                             "EmailAddress is: ": EmailAddress,
				                             "IsAdmin is: ": IsAdmin})
			  }
				  //err1 := db.QueryRow("select * from test.USER  where ID=?;",json.SeID) //db为sql.DB
			  //if err1 != sql.ErrNoRows {
			  // if err1,_:=db.Exec("select name,tel,e_mail_address from test.USER where ID = ?;",json.SeID);err1!=nil{
			   //c.JSON(http.StatusOK,json )}else {
			  // 	c.String(http.StatusNotFound,"failed")
			 //  }
			   }else {
			  c.JSON(http.StatusUnauthorized, gin.H{"status": "Unauthorized"})
		  }
	})
		// 监听并服务于 0.0.0.0:8080
	r.Run(":8080")
}

	func checkErr(err error) {
		if err != nil {
			fmt.Println("ID not exist")
			panic(err)
		}
	}