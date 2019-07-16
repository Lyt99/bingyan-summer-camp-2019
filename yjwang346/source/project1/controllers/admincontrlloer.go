package controllers

import (
	"crypto/aes"
	"crypto/cipher"
	"database/sql"
	"encoding/base64"
	"fmt"
	//"signin_Project/database"
	//jwt "github.com/appleboy/gin-jwt"
	"github.com/gin-gonic/gin"
	"net/http"
	"project1/models"

	_ "github.com/go-sql-driver/mysql"
)

func Checkadmin(c *gin.Context)(interface{}, error){
	//admin := c.Query("username")
	password := c.PostForm("password")

	if password=="123456"{
		c.String(200,"您以作为管理员登录")
	}else {
		c.String(400,"登录错误")
	}
	return nil,nil
}
//查询一条
func Getuserinfo(c *gin.Context) {
	//查询一条
	db, err := sql.Open("mysql", "root:hanru1314@tcp(localhost:3306)/bolgweb_gin?charset=utf8")
	Check(err)
	Checkadmin(c)//那岂不是每次都要输入密码？
	//获得要进行删除的用户的用户名
	delete_username := c.PostForm("delete_username")
	//所有的传参都是用小写的
	//s.Id = c.PostForm("id")
	//关键是传参的类型不对啊，还是最好就是使用id进行传参
	row := db.QueryRow("SELECT id,username,password,tel,email FROM userinfo WHERE username=?", delete_username)
	//这里SELECT后面选择一律采用小写格式
	var Id uint16
	var UserName string
	var Tel string
	var Psw string
	var Email string
	/*
	   row：Scan()-->将查询的结果从row取出
	       err对象
	       判断err是否为空，
	           为空，查询有结果，数据可以成功取出
	           不为空，没有数据，sql: no rows in result set
	*/
	error := row.Scan(&Id, &UserName, &Psw,&Tel, &Email)

	//fmt.Println(err)
	if error != nil {
		fmt.Println("查无数据。。")
	} else {
		fmt.Println(Id,UserName,AesDecrypt(Psw), Tel, Email)
	}
	db.Close()
}

//查询多条,但是这个方法比较麻烦

func getusersinfo() {
	/*
	   查询操作：
	*/
	//step2:打开数据库，建立连接
	db, _ := sql.Open("mysql", "root:0@tcp(localhost:3306)/bolgweb_gin?charset=utf8")

	//stpt3：查询数据库
	rows, err := db.Query("SELECT Id,UserName,Tel,Email FROM userinfo")
	if err != nil {
		fmt.Println("查询有误。。")
		return
	}
	//fmt.Println(rows.Columns()) //[uid username departname created]
	//创建slice，存入struct，
	datas := make([] models.SignForm, 0)
	//step4：操作结果集获取数据
	for rows.Next() {
		var Id int
		var UserName string
		var Tel string
		var Psw string
		var Email string
		if err := rows.Scan(&Id, &UserName, &Tel, &Psw,&Email); err != nil {
			fmt.Println("获取失败。。")
		}
		//每读取一行，创建一个user对象，存入datas2中
		user := models.SignForm{Id, Psw, UserName, Tel,Email}
		datas = append(datas, user)
	}
	//step5：关闭资源
	rows.Close()
	db.Close()

	for _, v := range datas {
		fmt.Println(v)
	}
}
//查询：处理查询后的结果：
//    思路一：创建结构体
//    思路二：将数据，存入到map中



//将数据库里的user全部输出,这个功能应该没有问题
//没有使用到gin.Context;
func Check(err error){     //因为要多次检查错误，所以干脆自己建立一个函数。
	if err!=nil{
		fmt.Println(err)
	}
}
func GetUsersinfo(c *gin.Context){
	db,err:=sql.Open("mysql","root:123456@tcp(127.0.0.1:3306)/employee")
	Check(err)
	Checkadmin(c)
	//query

	rows,err := db.Query("SELECT * FROM message")
	Check(err)

	c.String(http.StatusOK,"所有用户的信息如下：")
	for rows.Next(){
		var s models.SignForm
		err=rows.Scan(&s.Id,&s.UserName,&s.Psw,&s.Tel,&s.Email)
		Check(err)
		s.Psw = AesDecrypt(s.Psw)
		//似乎不用输出密码
		fmt.Println(s)
	}
	rows.Close()
}
//删除
func Deletuser(c *gin.Context){
	db,err:=sql.Open("mysql","root:123456@tcp(127.0.0.1:3306)/employee")
	Check(err)
	Checkadmin(c)

	delete_username := c.PostForm("username")
	results,err := db.Prepare("delete from userinfo where username = ?")
	Check(err)
	_, err = results.Exec(delete_username)
	db.Close()
}




//解密，将数据库中进行加密的密码转换成普通密码
func AesDecrypt(cryted string) string {
	// 转成字节数组
	key := "0123456789"
	crytedByte, _ := base64.StdEncoding.DecodeString(cryted)
	k := []byte(key)
	// 分组秘钥
	block, _ := aes.NewCipher(k)
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 加密模式
	blockMode := cipher.NewCBCDecrypter(block, k[:blockSize])
	// 创建数组
	orig := make([]byte, len(crytedByte))
	// 解密
	blockMode.CryptBlocks(orig, crytedByte)
	// 去补全码
	orig = PKCS7UnPadding(orig)
	return string(orig)
}
//去码
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}