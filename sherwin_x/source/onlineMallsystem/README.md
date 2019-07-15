## 实战商城系统：

基本功能：

- 用户分为买家和卖家

  - 卖家可以发布商品、删除商品
  - 买家可以购买商品

- 登录注册

  - 用户密码加密存储
  - 用户信息包括用户ID、密码（数据库中加密）、昵称、手机号、邮箱地址等，可自定义。

- 商品按照类别查询

  如：商品类别：电子设备、书籍资料、宿舍百货、美妆护肤、女装、男装、鞋帽配饰、门票卡券、其他

- 根据关键词搜索商品

  - 推荐用正则表达式

- 商品页面

  - 商品详细信息

    标题、简介、价格等

  - 图片

    图片可以存在本地，或者使用七牛云存储

- 个人信息页（类似于名片，自己和他人都可以查看）

  - 个人基本信息
  - 浏览量（个人信息页被他人访问次数，可考虑去重）

进阶功能：

- 图片压缩

  浏览时显示压缩的小图片，详细页显示大一点的图片

- 收藏夹(类似于购物车)

- 商品浏览量、收藏量等

- 热门查询、最新查询

  热门查询可在后台记录用户的浏览数据等信息

## 项目部署

### 1. 配置nginx

学习配置 nginx 做中间代理层，具体可从以下链接中选取部分学习，作为示例，夏令营之后可以好好研究，当然夏令营期间有时间也可以自行研究，遇到坑可以问我们。

[nginx 配置简介](https://juejin.im/post/5ad96864f265da0b8f62188f)

[openresty 实践](https://juejin.im/post/5aae659c6fb9a028d375308b)

### 2. 配置 docker

[Docker 从入门到实践](https://yeasy.gitbooks.io/docker_practice/content/install/ubuntu.html)

[Docker 实践](https://juejin.im/post/5b34f0ac51882574ec30afce)

### 3. 配置域名https (不要求)

前提：有已经备案的域名，有服务器

[Let's Encrypt 给网站加 HTTPS 完全指南](https://ksmx.me/letsencrypt-ssl-https/?utm_source=v2ex&utm_medium=forum&utm_campaign=20160529)