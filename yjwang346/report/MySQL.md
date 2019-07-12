链接：[Go操作MySQL](https://www.chaindesk.cn/witbook/17/263)

## 1.2连接数据库

要想使用go语言操作mysql，首先需要和mysql数据库建立连接，获取到DB对象。

导入包：

```go
import (
        　　"database/sql"
        　　_"github.com/Go-SQL-Driver/MySQL"
)
```

### 1.2.2  建立连接

使用Open函数：

Open函数：

```go
func Open(driverName, dataSourceName string) (*DB, error)

/*其中的两个参数：
    drvierName，这个名字其实就是数据库驱动注册到 database/sql 时所使用的名字.
        "mysql"
    dataSourceName，数据库连接信息，这个连接包含了数据库的用户名, 密码, 数据库主机以及需要连接的数据库名等信息.
        用户名:密码@协议(地址:端口)/数据库?参数=参数值
*/
db, err := sql.Open("mysql", "用户名:密码@tcp(IP:端口)/数据库?charset=utf8")
//例如：例如：db, err := sql.Open("mysql","root:111111@tcp(127.0.0.1:3306)/test?charset=utf8")
```

说明：

> 1. sql.Open并不会立即建立一个数据库的网络连接, 也不会对数据库链接参数的合法性做检验, 它仅仅是初始化一个sql.DB对象. 当真正进行第一次数据库查询操作时, 此时才会真正建立网络连接;
> 2. sql.DB表示操作数据库的抽象接口的对象，但不是所谓的数据库连接对象，sql.DB对象只有当需要使用时才会创建连接，如果想立即验证连接，需要用Ping()方法;
> 3. sql.Open返回的sql.DB对象是协程并发安全的.
> 4. sql.DB的设计就是用来作为长连接使用的。不要频繁Open, Close。比较好的做法是，为每个不同的datastore建一个DB对象，保持这些对象Open。如果需要短连接，那么把DB作为参数传入function，而不要在function中Open, Close。

### 1.3 DML操作 ：增删改

#### 1.3.1 操作增删改

有两种方法

1. 直接使用Exec函数添加

```go
func (db *DB) Exec(query string, args ...interface{}) (Result, error)
//示例代码：
//result, err := db.Exec("UPDATE userinfo SET username = ?,
//							departname = ? WHERE uid = ?", "王二狗","行政部",2)
```

2. 首先使用Prepare获得stmt，然后调用Exec添加。

建立连接后，通过操作DB对象的Prepare()方法，可以进行 

```go
func (db *DB) Prepare(query string) (*Stmt, error) {
    return db.PrepareContext(context.Background(), query)
}
//示例代码
/*
stmt,err:=db.Prepare("INSERT INTO userinfo(username,departname,created) values(?,?,?)") 
补充完整sql语句，并执行
result,err:=stmt.Exec("韩茹","技术部","2018-11-21")
*/
```