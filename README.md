# 冰岩作坊程序组2019夏令营

***欢迎参加冰岩作坊夏令营！***



## 前言

- 请先Fork(右上角)此仓库，本次夏令营要求代码、日报等全部托管在你们fork后的github仓库中
- 日报和周报等不需要写太多，只需要介绍每天学习了什么，以及适当记录你认为的重点即可
- 之后的代码检查(code review)，采用pull request(PR)的形式

## 操作说明

- fork 此仓库
- 在你的仓库操作的时候请不要对他人目录进行任何操作
- 你的操作权限仅限于你的目录，目录名字为你的 github ID，若仓库中没有你的目录请自行创建
- 提交 PR 的时候自行查看是否存在代码冲突，如果存在自行解决之后再提交 PR
- 提交 PR 是提交到 dev 分支，不是 master 分支
- 目录结构推荐如下：
  - reports文件夹 - 日报
  - source文件夹 - 源码，各项目创建不同的文件夹

## 学习安排

### 1.语言

- Golang语法
  - [官方链接](https://golang.org/)
  - [官方中文教程](https://tour.go-zh.org/welcome/1)
  - [语言规范](https://go-zh.org/ref/spec)

- 书籍推荐
  - [《The Go Programming Language》中文版](https://www.gitbook.com/book/yar999/gopl-zh/details)
  - [《Effective Go》中英双语版](https://www.gitbook.com/book/bingohuang/effective-go-zh-en/details)
  - [Go语言实战](http://download.csdn.net/download/truthurt/9858317)
  - [Go Web编程](https://wizardforcel.gitbooks.io/build-web-application-with-golang/content/index.html)可以了解基本web开发，比较推荐入门

### 2.框架

> 在学习框架的过程中，了解一下MVC架构，并在热身项目中加以应用。

- gin（推荐）

  - [gin英文文档](https://github.com/gin-gonic/gin)

  - [ Gin 文档中文翻译](https://learnku.com/docs/gin-gonic/2018/gin-readme/3819)

- beego

  - [beego: 简约 & 强大并存的 Go 应用框架](https://beego.me/docs/intro/)

- Iris

  - [Iris英文文档](https://github.com/kataras/iris)

  - [Iris文档中文翻译](https://studyiris.com/doc/)

- echo(可以以后自学)

  - [echo英文文档](https://echo.labstack.com/guide)
  - [echo文档中文翻译](http://go-echo.org/)

- 其他

### 3. HTTP相关

- HTTP请求方法：GET、POST、PUT、UPDATE等

- HTTP状态码：404、200、400、401、301、500等

- HTTP数据传输格式：[json](https://www.runoob.com/json/json-syntax.html)、form表单

- HTTP报文格式（大致了解就行、不用深入学习）

- 前后端如何交互？前后端分离是什么？

  前端如何获取后端返回的数据，如何发送请求，后端如何根据前端发过来的请求，回应请求，如何辨别不同的请求

### 4. 数据库相关

- MySQL（推荐优先学习）
- MongoDB（后期推荐学习、可以在夏令营之后研究，有能力的可以夏令营用，和go搭配比较好用）
- Redis（基于内存的非关系型数据库）

### 5. 其他知识

**认证：**

熟悉以下三种前后端认证方式，一般在登录时使用

- cookie
- session
- JWT

**加密算法：**

- 对称加密
- 非对称加密
- 哈希算法

### 6. 相关工具

- 编辑器：goland、vscode

- 后台接口测试工具：postman

  

## 热身项目

**成员管理系统**

实现内容：

- 管理员和普通用户

- 用户注册和登录

  用户信息包括用户ID、密码（数据库中加密）、昵称、手机号、邮箱地址

- 管理员
  - 删除普通用户
  - 获取一个成员、所有成员信息

- 普通用户
  - 更改个人信息