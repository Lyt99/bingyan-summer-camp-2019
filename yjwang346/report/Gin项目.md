![gin_1](http://image.chaindesk.cn/gin_2.png/mark)

执行原理：

- 1. 进入main.go,初始化路由，以及端口号
- 1. 根据浏览器输入的URL地址，在router路由器中找到对应的路由函数方法
- 1. 根据路由中URL后指定的函数，在Controller中找到对应的方法函数
- 1. Controller中调用models关于数据库方面的方法函数
- 1. 渲染html页面，执行js,css等效果