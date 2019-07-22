## 基于gin的商城后台

### 功能

在为前端提供用户的登录注册接口的基础上，基于规定API文档设计了用户登录后的个人资料、商品管理、收藏、浏览等权限操作

### 环境

系统：Manjaro gnome 

语言环境：Gin-基于golang的web框架

数据库：MongoDB

编译器：goland

测试工具：postman

### 开始

1. 初始化MongoDB服务

```bash
$ source ~/.bashrc
$ mongod --dbpath /home/sherwin/tools/mongodb/data
$ mongo
```

2. 在终端中打开文件夹并运行

```bash
$ cd path/to/project
$ go run main.go
```

3. 通过postman进行接口测试

### 接口介绍

**注册**：

- 方法:POST

**登录**：

**查看用户资料**

**查看我的资料**

**修改我的资料**

**查看我的发布**

**查看我的收藏**

**收藏某个商品**

**删除某个收藏**

**获取商品列表**

**获取热搜词**

**发布新商品**

**查看商品详情**

**删除某个商品**