package database

import (
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"fmt"
	//"log"
	"project1/models"
	"signin_Project/database"
)


func signup(){
	// step2：打开数据库，相当于和数据库建立连接：db对象
	db,err:=sql.Open("mysql","root:0@tcp(localhost:3306)/bolgweb_gin?charset=utf8")
	if err !=nil{
		fmt.Println("连接失败。。")
		return
	}
	//step3：插入一条数据

	stmt,err:=db.Prepare("INSERT INTO userinfo(username,departname,created) values(?,?,?)")
	if err !=nil{
		fmt.Println("操作失败。。")
	}
	//补充完整sql语句，并执行
	result,err:=stmt.Exec("韩茹","技术部","2018-11-21")
	if err !=nil{
		fmt.Println("插入数据失败。。")
	}
	//step4：处理sql操作后的结果
	lastInsertId,err:=result.LastInsertId()
	rowsAffected,err:=result.RowsAffected()
	fmt.Println("lastInsertId",lastInsertId)
	fmt.Println("影响的行数：", rowsAffected)
	//step5：关闭资源
	stmt.Close()
	db.Close()

}

//插入
func InsertUser(user models.SignForm) (int64, error) {
	return database.ModifyDB("insert into users(Id,UserName,Psw,Tel,Email) values (?,?,?,?)",
		user.Id,user.UserName, user.Psw, user.Tel, user.Email)
}

//按条件查询
func QueryUserWightCon(con string) int {
	sql := fmt.Sprintf("select id from users %s", con)
	fmt.Println(sql)
	row := database.QueryRowDB(sql)
	id := 0
	row.Scan(&id)
	return id
}

//根据用户名查询id
func QueryUserWithUsername(username string) int {
	sql := fmt.Sprintf("where username='%s'", username)
	return QueryUserWightCon(sql)
}

/*
//根据用户名和密码，查询id
func QueryUserWithParam(username ,password string)int{
    sql:=fmt.Sprintf("where username='%s' and password='%s'",username,password)
    return QueryUserWightCon(sql)
}

*/
func QueryUser() {
	/*
	   查询一条
	*/
	db, _ := sql.Open("mysql", "root:hanru1314@tcp(127.0.0.1:3306)/blogweb_gin")

	row := db.QueryRow("SELECT uid,username,departname,created FROM userinfo WHERE uid=?", 1)
	var uid int
	var username, departname, created string
	/*
	   row：Scan()-->将查询的结果从row取出
	       err对象
	       判断err是否为空，
	           为空，查询有结果，数据可以成功取出
	           不为空，没有数据，sql: no rows in result set
	*/
	err := row.Scan(&uid, &username, &departname, &created)

	//fmt.Println(err)
	if err != nil {
		fmt.Println("查无数据。。")
	} else {
		fmt.Println(uid, username, departname, created)
	}
	db.Close()

}