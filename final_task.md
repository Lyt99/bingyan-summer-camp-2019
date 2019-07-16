# 冰岩2019夏令营最终任务 API文档

## 数据交换格式

数据(request&response)使用json，response格式如下

```json
{
    "success": true,
    "error": "错误信息",
    "data": ""
}
```



success: 布尔型，请求是否成功(http状态码200)

error: 字符串，错误信息，当请求不成功的时候，用于提示调用者错误原因。成功时可留空

data: 成功后返回的数据，可以是任意类型，失败的时候留空



## 用户身份认证约定

使用jwt(JSON Web Token)进行认证， 将token放入http头的Authorization字段，并使用Bearer开头，如

token为**eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c**

HTTP头为

Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c

## 商品类别

商品类别使用int表示，规定

- 1 电子设备
- 2 书籍资料
- 3 宿舍百货
- 4 美妆护肤
- 5 女装
- 6 男装
- 7 鞋帽配饰
- 8 门票卡券
- 9 其它



## API列表

### POST /user

**注册用户**

#### request

```json
{
    "username": "用户名",
    "password": "密码",
    "nickname": "昵称",
    "mobile": "手机号",
    "email": "邮箱"
}
```

#### response

##### 成功

```json
{
    "success": true,
    "error": "",
    "data": "ok"
}
```

##### 失败

用户名已存在、手机号，邮箱不合法等，例如

```json
{
    "success": false,
    "error": "用户名已存在",
    "data": null
}
```



### POST /user/login

用户登录，并获取token

#### request

```json
{
    "username": "用户名",
    "password": "密码"
}
```

password: 用户密码，可以事先加密(不能和存入数据库的加密方式一样！)，也可以明文

#### response

##### 成功

```json
{
    "success": true,
    "error": "",
    "data": "jwt token"
}
```



### GET /commodities

**获得商品列表**

#### request

```json
{
    "page": 0,
    "limit": 20,
    "category": 0,
    "keyword":""
}
```

page: 第几页，从0开始

limit: 每页的商品数量

category: 商品分类，为0即获得全部分类商品

keyword:搜索关键词，为空则搜索全部商品

#### response

##### 成功

```json
{
    "success": true,
    "error": "",
    "data": [
        {
            "id": "商品ID",
            "title": "商品标题",
            "price": 1.0,
            "category": 1,
            "picture": "https://blablabla.bla/bla.jpg"
        }
    ]
}
```

**商品按照时间降序排序**

其中，category为商品的类别，picture为商品图片的url

商品ID为商品插入数据库时分配的一个独一无二的ID，可以使用MongoDB的**_id**字段或者自增id

这里不显示商品的简介，在详细页面显示

##### 失败

***没有商品不算失败！返回空列表！***

参数不合法(为负数)等



### GET /commodities/hot

获得热门查询关键词

#### request

无

#### response

```json
{
    "success": true,
    "error": "",
    "data": ["关键词1", "关键词2"]
}
```



### POST /commodities

发布商品

#### request

```json
{
    "title": "商品标题",
    "desc": "商品简介",
    "category": 1,
    "price": 1.0,
    "picture": "https://lablabla.bla/bla.jpg"
}
```



#### response

##### 成功

```json
{
    "success": true,
    "error": "",
    "data": "ok"
}
```

##### 失败

输入信息不合法等



### GET /commodity/:id

获得商品详细信息

id : 商品ID，如 /commodity/123

#### request

无

#### response

##### 成功

```json
{
    "success": true,
    "error": "",
    "data": {
        "pub_user": "商品发布者用户ID",
        "title": "商品标题",
        "desc": "商品简介",
        "category": 1,
        "price": 1.0,
        "picture": "https://lablabla.bla/bla.jpg",
        "view_count": 20,
        "collect_count": 0
	}
}
```



pub_user: 用户id，可以直接使用用户注册时的ID

view_count: 浏览数

collect_count: 收藏数

##### 失败

商品不存在等



### DELETE /commodity/:id

删除商品

id: 商品ID

#### request

无

#### response

##### 成功

```json
{
    "success": true,
    "error": "",
    "data": "ok"
}
```

***不能让用户删除不是自己的商品！***



##### 失败

商品不存在、不是自己的商品等

### GET /me

获得自己的个人信息

#### request

无

#### response

#####  成功

```json
{
    "success": true,
    "error": "",
    "data": {
        "username": "用户名",
        "nickname": "昵称",
        "mobile": "手机号",
        "email": "邮箱",
        "total_view_count": 10,
        "total_collect_count": 0
	}
}
```



total_view_ count: 商品总浏览量

total_collect_count: 商品总收藏量



### POST /me

修改我的个人信息

#### request

```json
{
    "username": "用户名",
    "password": "密码",
    "nickname": "昵称",
    "mobile": "手机号",
    "email": "邮箱"
}
```

**前端应保持不作修改的值为原值(密码除外)**

**密码为空表示不修改密码**

#### response

##### 成功

```json
{
    "success": true,
    "error": "",
    "data": "ok"
}
```

##### 失败

手机号、邮箱不合法等



### GET /me/commodities

获得自己发布的商品

#### request

无

#### response

##### 成功

```json
{
    "success": true,
    "error": "",
    "data": [
        {
            "id": "商品ID",
            "title": "商品标题"
        }
     ]
}
```

##### 失败

***没有商品不算失败！返回空列表！***

未登录等

### 

### GET /me/collections

获得我收藏的商品

#### request

无

#### response

##### 成功

```json
{
    "success": true,
    "error": "",
    "data": [
        {
            "id": "商品ID",
            "title": "商品标题"
        }
     ]
}
```

#####  失败

***没有商品不算失败！返回空列表！***

未登录等



### POST /me/collections

收藏商品

#### **request**

```json
{
    "id": "商品ID"
}
```

#### response

##### 成功

```json

    "success": true,
    "error": "",
    "data": "ok"
}
```

##### 失败

未登录、商品不存在等



### DELETE /me/collections

删除商品

#### request

```json
{
    "id": "商品ID"
}
```

#### response

##### 成功

```json
{
    "success": true,
    "error": "",
    "data": "ok"
}
```

##### 失败

商品不存在于收藏中等



### GET /user/:id

查看其它用户信息

#### request

无

#### response

##### 成功

**如果调用该接口的是自己，那么该接口等效于 GET /me**

```json
{
    "success": true,
    "error": "",
    "data": {
        "nickname": "昵称",
        "email": "邮箱",
        "total_view_count": 10,
        "total_collect_count": 0
	}
}
```

##### 失败

用户不存在等





