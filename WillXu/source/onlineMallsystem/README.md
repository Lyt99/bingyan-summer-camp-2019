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

1. 拉取项目至本地

```bash
 $ git clone
```


2. 初始化MongoDB服务

```bash
$ source ~/.bashrc
$ mongod --dbpath /home/sherwin/tools/mongodb/data
```

3. 在终端中打开文件夹并运行

```bash
$ cd path/to/project
$ go run main.go
```

4. 通过postman进行接口测试
