### 1. gin框架介绍

gin框架是Go语言进行web开发（api开发，微服务开发）框架中，功能和Martini框架类似的API，但是性能却特别好的一个框架（比Martini快将近40倍吧），所以如果你特别在乎性能，那么Gin会是一个比较好的选择。

gin框架主要基于httprouter模块进行实现。gin框架和httprouter都是一个开源的框架。

微服务本身即是一种开发模式，将业务拆分成为一个个细小的微服务模块，然后以api（rpc）方式对外提供实现，实现的功能是一个独立的业务模块，那么使用轻量级的gin便是一个不错的选择。

### 2. gin框架包含的功能模块

gin框架包含了以下主要功能：

> http请求的Context上下文 
>
> 基础的auth认证模块 带颜色的logger模块 
>
> 运行模式mode设定 
>
> 响应处理的responsewriter模块
>
> 以及路由组routergroup





[Gin-Web框架](https://www.cnblogs.com/tudaogaoyang/p/8056249.html)

> 执行的原理：
>
> 1->进入main.go,初始化路由，以及端口号
>
> 2->根据浏览器输入的URL地址，在router路由器中找到对应的路由函数方法
>
> 3->根据路由中URL后指定的函数，在Api/Controller中找到对应的方法函数 
>
> 4->Api/Controller调用关于数据库方面的方法函数
>
> 5->展示html页面
>
> 6->渲染html页面，执行js,css等效果



ps：我下载了vendor工具到本地，但是没有进行后续操作

## gin框架学习



```go
package main
import (     
"github.com/gin-gonic/gin"  
"net/http" 
) 
func main()      
   router := gin.Default() 
   router.GET("/", func(c *gin.Context) {         
      c.String(http.StatusOK, "Hello World")    
   })     
   router.Run(":8000")
}
//然后在浏览器里面输入 localhost:8000就可以输出Hello World在浏览器的界面里显示
//还可以直接在postman里面GET localhost:8000		一样能得到
```

#### restful路由、query string

\<c:param>为指定URL发送两个参数

```go
func main() {
	router := gin.Default()

	// 这个处理器可以匹配 /user/john ， 但是它不会匹配 /user
	router.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	// 但是，这个可以匹配 /user/john 和 /user/john/send
	// 如果没有其他的路由匹配 /user/john ， 它将重定向到 /user/john/
	router.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		c.String(http.StatusOK, message)
	})

	router.Run(":8080")
}
//在postman里面输入localhost:8080/user/wangyj,也就是将name用实际上的值来代替
//输出Hello wangyj
//GET localhost:8080/user/wangyj/asdasd
//输出wangyj is /asdasd
```

**/:name**这一栏里面，代码部分冒号  **:**  不可以省；冒号 :   是为了把name和user区分开来



```go
func main() {
	router := gin.Default()

	// 查询字符串参数使用现有的底层 request 对象解析。
	// 请求响应匹配的 URL： /welcome?firstname=Jane&lastname=Doe
	router.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		// 这个是 c.Request.URL.Query().Get("lastname") 的快捷方式。
		lastname := c.Query("lastname")

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})
	router.Run(":8080")
}
//这种形式的GET在postman里面应该怎样操作？
//另外，一个main里面有多个 router.GET 根目录不一样执行的函数也不一样
```

先把gin的这些比如c.DefaultQuery、c.Query什么的都弄明白



MVC，相当于把上面的这些代码封装出来，

比如上面的GET中的func可以单另拿出来构成一个**welcom**函数，然后再调用它，就构成了control层，然后底层的一些逻辑就是module层……





Goland里面要习惯用命令行，运行时候就是Terminal栏里面输入"go run"+"exam.go" (文件名)

