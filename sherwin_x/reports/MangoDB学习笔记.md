
### 文档

MangoDB的核心概念之一

形式：{"键":"键对应的键值","第二个键":"第二个键对应的键值",.......}

键：除了某些保留字符意外的**所有通用字符串**

键值：任意UTF-8字符

### 集合

文档对应数据库中的行，集合对应数据库中的表

注意，这里的集合是**无模式**的，意味着自由度很大

### 数据库

多个文档组成集合，多个集合组成数据库

### MangoDB客户端

### 数据类型

| 数据类型           | 描述                                                         |
| :----------------- | :----------------------------------------------------------- |
| String             | 字符串。存储数据常用的数据类型。在 MongoDB 中，UTF-8 编码的字符串才是合法的。 |
| Integer            | 整型数值。用于存储数值。根据你所采用的服务器，可分为 32 位或 64 位。 |
| Boolean            | 布尔值。用于存储布尔值（真/假）。                            |
| Double             | 双精度浮点值。用于存储浮点值。                               |
| Min/Max keys       | 将一个值与 BSON（二进制的 JSON）元素的最低值和最高值相对比。 |
| Array              | 用于将数组或列表或多个值存储为一个键。                       |
| Timestamp          | 时间戳。记录文档修改或添加的具体时间。                       |
| Object             | 用于内嵌文档。                                               |
| Null               | 用于创建空值。                                               |
| Symbol             | 符号。该数据类型基本上等同于字符串类型，但不同的是，它一般用于采用特殊符号类型的语言。 |
| Date               | 日期时间。用 UNIX 时间格式来存储当前日期或时间。你可以指定自己的日期时间：创建 Date 对象，传入年月日信息。 |
| Object ID          | 对象 ID。用于创建文档的 ID。                                 |
| Binary Data        | 二进制数据。用于存储二进制数据。                             |
| Code               | 代码类型。用于在文档中存储 JavaScript 代码。                 |
| Regular expression | 正则表达式类型。用于存储正则表达式。                         |

### 基本用法

一些Mongo shell的基本命令及用法

    db.help()               	 //显示数据库操作命令
    db.foo.help()                //显示集合foo的操作命令
    
    db             			    //展示所处数据库
    db.dropDatabase()			//删除数据库（进入再执行）
    db.foo.insert()      	     //创建名为foo的集合（表）并插入文档（数据）	             
    db.foo.remove({})            //删除foo集合内的所有文档（保留foo集合）
    db.foo.drop()				//删除集合foo
    
    db.foo.find()                //对于当前数据库中的foo集合进行数据查找（由于没有条件，会列出所有数据）    db.foo.find( { a : 1 } )     //对于当前数据库中的foo集合进行条件查找，条件是数据中有一个键为a，键值为1
    
    use <db_name>                //转到指定数据库，不存在就创建
    
    show dbs                     //展示所有的数据库
    show collections/tables	     //展示所有的集合
    show users                   show users in current database
    show profile                 show most recent system.profile entries with time >= 1ms
    show logs                    show the accessible logger names
    show log [name]              prints out the last segment of log in memory, 'global' is default


​    
    exit                         quit the mongo shell
#### 数据库的增删查

```
use db_name					//新增（已存在则切换）数据库
db.dropDatabase()		     //删除当前数据库
show dbs				    //查询显示所有数据库
```

#### 集合的增删查

```
db.foo.insert({"a":"1"})	 //新增名为foo的集合并插入文档
db.foo.drop()			    //删除当前数据库内名为foo的集合
show tables				    //查询显示当前数据库内所有的集合
```

#### 文档的增删查

```
doc={
	"name":"joe",
	"addr":"hina"
	"tel":"123125126"
}
db.foo.insert(doc)	         //新增单条文档doc至集合foo中

db.foo.remove({})     		 //删除foo中的所有文档
db.foo.remove({"a":"1"})	 //删除foo中的指定文档

db.foo.find()				//查询集合foo中的所有文档
db.foo.find({"a","1"})   	 //指定条件查询集合foo中的某个文档
```

<!--20190705-->

