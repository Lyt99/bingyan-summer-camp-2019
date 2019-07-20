## 实战商城系统：

TO-DO：

- [x] **注册**：POST /user 

- [x] **登录**：POST /user/login

- [x] **查看某位用户资料**：GET /user/:id

  

- [x] **查看个人资料**：GET /me
- [ ] **修改个人资料**：POST /me：　　　　　　　待完善
- [x] **查看我的发布**：GET /me/commodities：　 待完善
- [ ] **查看我的收藏**：GET /me/collections
- [ ] **收藏某个商品**：POST /me/collections
- [ ] **删除某个收藏**：DELETE /me/collections



- [ ] **获取商品列表**：GET /commodities
- [ ] **获取热搜词**：GET /commodities/hot
- [x] **发布新商品**：POST /commodities



- [x] **某个商品详情**：GET /commodity/:id
- [x] **删除某个商品**：DELETE /commodity/:id



基本功能：

- 用户可以买卖商品

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