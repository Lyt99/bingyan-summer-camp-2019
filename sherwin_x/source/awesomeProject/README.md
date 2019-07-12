**成员管理系统**

- 管理员和普通用户

- 用户注册和登录

  用户信息包括用户ID、密码（数据库中加密）、昵称、手机号、邮箱地址

- 管理员

  - 删除普通用户
  - 获取一个成员、所有成员信息

- 普通用户

  - 更改个人信息

**环境配置**

系统：[Manjaro gnome](https://manjaro.org/download/) 

语言环境：[Go](https://golang.org/)

编译器：goland

测试工具：postman

**开始**

连接MongoDB数据库(假设数据库路径在`/home/sherwin/tools/mongodb/data`)

```bash
$ source ~/.bashrc
$ mongod --dbpath /home/sherwin/tools/mongodb/data
```

编译器中运行，postman测试



