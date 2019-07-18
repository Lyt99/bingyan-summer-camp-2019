# mongoDB
> 我希望今天我用得上它

mongoDB使用BSON格式存储数据
从大到小数据结构依次为
数据库  database

集合    collection

文档    document

域      field

索引    index

主键    primary key

> 注意：文档中的键值对是有序的，区分大小写，不可以有重复的键

集合拥有元数据，（我觉得很像静态变量）元数据位于<dbname>.system.*命名空间下，包含了多种特殊数据
比如

![1563005871282](C:\Users\Sid\AppData\Roaming\Typora\typora-user-images\1563005871282.png)

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

### ObjectId

ObjectId与主键类似，生成速度快，便于排序。

有十二个字节

分别为：

![1563006236174](C:\Users\Sid\AppData\Roaming\Typora\typora-user-images\1563006236174.png)

每一个文档中都要有一个**_Id**键,这个值可以为任何类型, 默认为一个ObjectId对象, 由于其中包含了时间戳所以无需在加上时间戳字段, 可以通过getTimeStamp函数来获取文档的创建时间.

BSON包含了一个时间戳,用于MongoDB内部使用, 与普通日期类型无关. 时间戳有64位, 前32位为一个time_t,

后32位为一个在秒操作中递增的序数. 

## 创建数据库

与git创建分支类似, 创建数据库使用**use DATABASE_NAME**的方式, 如果数据库不存在则创建, 存在就会切换过去

查看数据库使用**show dbs**

默认数据库为**test**, 如果没有建立新的数据库,集合会存放在test数据库中

> 集合需要在插入内容后才会创建

## 删除数据库与集合

删除数据库的语句为**db.dropDatabase()**, 这样就删除了当前数据库. 

删除集合的语句为**db.<Collection_Name>.drop()**, 这样就删除了<Collection_Name>集合

## 创建集合

创建集合的语句为**db.createCollecion(\<Name\>, options)**

> \<Name\>: 创建的集合名称, options: 可选参数(指定有关内存大小以及索引的选项)
>
> 参数字段如下

| 字段        | 类型 | 描述                                                         |
| :---------- | :--- | :----------------------------------------------------------- |
| capped      | 布尔 | （可选）如果为 true，则创建固定集合。固定集合是指有着固定大小的集合，当达到最大值时，它会自动覆盖最早的文档。 **当该值为 true 时，必须指定 size 参数。** |
| autoIndexId | 布尔 | （可选）如为 true，自动在 _id 字段创建索引。默认为 false。   |
| size        | 数值 | （可选）为固定集合指定一个最大值（以字节计）。 **如果 capped 为 true，也需要指定该字段。** |
| max         | 数值 | （可选）指定固定集合中包含文档的最大数量。                   |

## 插入文档

MongoDB使用BSON格式, 与JSON格式基本一致, MongoDB使用 insert() 或者 save() 方法向集合中插入文档

```shell
db.COLLECTION_NAME.insert(document)
```



当然document可以是一个变量

## 更新文档

使用**update()**和**save()**方法来更新集合中的文档.

```shell
db.collection.update(
   <query>,
   <update>,
   {
     upsert: <boolean>,
     multi: <boolean>,
     writeConcern: <document>
   }
)
```

方法用于更新已经存在的文档



>  参数说明
>
> - **query** : update的查询条件，类似sql update查询内where后面的。
> - **update** : update的对象和一些更新的操作符（如$,$inc...）等，也可以理解为sql update查询内set后面的
> - **upsert** : 可选，这个参数的意思是，如果不存在update的记录，是否插入objNew,true为插入，默认是false，不插入。
> - **multi** : 可选，mongodb 默认是false,只更新找到的第一条记录，如果这个参数为true,就把按条件查出来多条记录全部更新。
> - **writeConcern** :可选，抛出异常的级别。



```shell
db.collection.save(
   <document>,
   {
     writeConcern: <document>
   }
)
```

方法可以通过保存新的文档替换原来的文档

## 删除文档

我们使用**remove()**方法删除文档

方法基本格式如下

```shell
db.collection.remove(
   <query>,
   {
     justOne: <boolean>,
     writeConcern: <document>
   }
)
```