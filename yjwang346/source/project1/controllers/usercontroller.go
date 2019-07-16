package controllers

import (
	"fmt"

	"net/http"
	"database/sql"
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"

	"project1/models"
	"github.com/gin-gonic/gin"
)

var idnum int = 0
func Signup(c *gin.Context){
	db,err:=sql.Open("mysql","root:0@tcp(localhost:3306)/mytest?charset=utf8")

	defer db.Close()
	Check(err)
	idnum = idnumber(idnum,db)
	idnum++				//计数完之后还要进行一次+
	//但问题是，程序关闭之后，id也会变成零该怎么办？
	var newuser models.SignForm
	newuser.UserName = c.PostForm("username")
	newuser.Psw = c.PostForm("password")
	password := AesEncrypt(newuser.Psw)
	//进行加密
	newuser.Tel = c.PostForm("telephone")
	newuser.Email = c.PostForm("email")
	//密码换成了进行加密之后得到的密码
	CreateTableWithUser(db,idnum ,newuser.UserName,password,newuser.Tel, newuser.Email)
	//加入数据库内容？
	checksignup(db,idnum)
	c.String(http.StatusOK,"注册成功！")
}
//向数据库中增加用户相关信息
func CreateTableWithUser(db *sql.DB,id int,username,password, tel, email string) {
	stmt, err := db.Prepare("INSERT userinfo SET id=?,username=?,password=?, tel=?, email=?")
	Check(err)

	_, err = stmt.Exec(id,username,password, tel, email)
	Check(err)
}

//比较看用户是否已经注册
func checksignup(db *sql.DB,id int) {
		//查询一条

		//db, _ := sql.Open("mysql", "root:hanru1314@tcp(localhost:3306)/bolgweb_gin?charset=utf8")
		row := db.QueryRow("SELECT Id,UserName,Tel,Email FROM userinfo WHERE uid=?", id)
		//var uid int
		//var username, departname, created string
		var check models.SignForm
		 /*  row：Scan()-->将查询的结果从row取出
	 	      err对象,判断err是否为空，
		           为空，查询有结果，数据可以成功取出
	 	          不为空，没有数据，sql: no rows in result set
		*/
		err := row.Scan(&check.Id, &check.UserName)

		//fmt.Println(err)
		if err != nil {
			fmt.Println("该用户名可以使用")
		} else {
			fmt.Println("该用户名已经使用，请更换用户名")
		}
}


//方法很笨，	用来得到数据库中有多少数据,得到注册时用户的号码数
func idnumber(idnum int,db *sql.DB) int {
	rows,err := db.Query("SELECT * FROM message")
	Check(err)
	for rows.Next(){
		idnum++
	}
	return idnum
}


func LoginPost(c *gin.Context) (interface{}, error){
	//获取表单信息
	db, err := sql.Open("mysql", "root:hanru1314@tcp(localhost:3306)/bolgweb_gin?charset=utf8")
	Check(err)
	defer db.Close()
	username := c.PostForm("username")
	password := c.PostForm("password")
	fmt.Println("username:", username, ",password:", password)

	id := QueryUserWithParam(db,username, AesEncrypt(password))
	fmt.Println("id:", id)
	if id > 0 {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "登录成功"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "登录失败"})
	}
	return nil,nil
}
//根据用户名和密码，查询id
func QueryUserWithParam(db *sql.DB,username ,password string)int{
	sql:=fmt.Sprintf("where username='%s' and password='%s'",username,password)
	return QueryUserWightCon(db,sql)
}
//按条件查询
func QueryUserWightCon(db *sql.DB,con string) int {
	sql := fmt.Sprintf("select id from users %s", con)
	fmt.Println(sql)
	row := QueryRowDB(db,sql)
	id := 0
	row.Scan(&id)
	return id
}
func QueryRowDB(db *sql.DB,sql string) *sql.Row{
	return db.QueryRow(sql)
}

func Changeinfo(c *gin.Context){
	LoginPost(c)		//通过用户名来更改数据库相关内容

	username := c.PostForm("username")
	new_name := c.PostForm("new_username")
	new_password := c.PostForm("new_password")
	new_tel := c.PostForm("new_tel")
	new_email := c.PostForm("new_email")


	Changepassword(username,new_password)
	Changetel(username,new_tel)
	Changeemail(username,new_email)
	Changename(&username,new_name)
}
//这里比较特别，换了名字之后，那么如果后面还要进行操作，那么传递的名字也要变
func Changename(username *string,new_name string){
	if new_name=="" {
		return
	}
	db, err := sql.Open("mysql", "root:hanru1314@tcp(localhost:3306)/bolgweb_gin?charset=utf8")
	Check(err)
	defer db.Close()

	stmt,erro:=db.Prepare("UPDATE userinfo SET username=? WHERE username=?")
	Check(erro)
	_, err = stmt.Exec(new_name, username)
	Check(err)
}

func Changepassword(username,new_password string){
	if new_password=="" {
		return
	}
	db, err := sql.Open("mysql", "root:hanru1314@tcp(localhost:3306)/bolgweb_gin?charset=utf8")
	Check(err)
	defer db.Close()
	stmt,erro:=db.Prepare("UPDATE userinfo SET password=? WHERE username=?")
	Check(erro)
	_, err = stmt.Exec(new_password, username)
	Check(err)
}

func Changetel(username,new_tel string){
	if new_tel=="" {
		return
	}
	db, err := sql.Open("mysql", "root:hanru1314@tcp(localhost:3306)/bolgweb_gin?charset=utf8")
	Check(err)
	defer db.Close()
	stmt,erro:=db.Prepare("UPDATE userinfo SET tel=? WHERE username=?")
	Check(erro)
	_, err = stmt.Exec(new_tel, username)
	Check(err)
}


func Changeemail(username,new_email string){
	if new_email=="" {
		return
	}
	db, err := sql.Open("mysql", "root:hanru1314@tcp(localhost:3306)/bolgweb_gin?charset=utf8")
	Check(err)
	defer db.Close()
	stmt,erro:=db.Prepare("UPDATE userinfo SET email=? WHERE username=?")
	Check(erro)
	_, err = stmt.Exec(new_email, username)
	Check(err)
}


//CBC 模式加密, 省略了传递参数里面的key string，为了方便，直接在这个函数体里面定义了key
func AesEncrypt(orig string) string {
	// 转成字节数组
	key := "0123456789"
	origData := []byte(orig)
	k := []byte(key)
	// 分组秘钥
	// NewCipher该函数限制了输入k的长度必须为16, 24或者32
	block, _ := aes.NewCipher(k)
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 补全码
	origData = PKCS7Padding(origData, blockSize)
	// 加密模式
	blockMode := cipher.NewCBCEncrypter(block, k[:blockSize])
	// 创建数组
	cryted := make([]byte, len(origData))
	// 加密
	blockMode.CryptBlocks(cryted, origData)
	return base64.StdEncoding.EncodeToString(cryted)
}
//补码
//AES加密数据块分组长度必须为128bit(byte[16])，密钥长度可以是128bit(byte[16])、192bit(byte[24])、256bit(byte[32])中的任意一个。
func PKCS7Padding(ciphertext []byte, blocksize int) []byte {
	padding := blocksize - len(ciphertext)%blocksize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
