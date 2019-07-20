参考链接: [GIn框架开发与实战](https://chaindesk.cn/witbook/19/330)

### Page 1

- 1、`router:=gin.Default()`：这是默认的服务器。使用gin的`Default`方法创建一个路由`Handler`；
- 2、然后通过Http方法绑定路由规则和路由函数。不同于`net/http`库的路由函数，gin进行了封装，把`reques`t和`response`都封装到了`gin.Context`的上下文环境中。
- 3、最后启动路由的Run方法监听端口。还可以用`http.ListenAndServe(":8080", router)`，或者自定义Http服务器配置。

### Page 2 Router

除了默认服务器中 `router.Run()` 的方式外，还可以用 `http.ListenAndServe()`，比如

```go
func main() {
    router := gin.Default()
    http.ListenAndServe(":8080", router)
}
```

api参数通过Context的Param方法来获取，这里的Context就是指 func(c *gin.Context)

冒号`:`加上一个参数名组成路由参数。可以使用c.Params的方法读取其值。当然这个值是字串string。

**2.2.2 URL参数**

URL 参数通过 DefaultQuery 或 Query 方法获取。***c.DefaultQuery***方法读取参数，其中当参数不存在的时候，提供一个默认值。使用***Query方法***读取正常参数，当参数不存在的时候，返回空字串。

**2.2.3 表单参数**

http的报文体传输数据就比query string稍微复杂一点，常见的格式就有四种

默认情况下，c.PostFROM解析的是`x-www-form-urlencoded`或`from-data`的参数。

使用PostForm形式,注意必须要设置Post的**type**(HTML中的type)，同时此方法中忽略URL中带的参数,所有的参数需要从Body中获得。

**2.2.4 文件上传**

`multipart/form-data`转用于文件上传

```go
func main() {
    router := gin.Default()
    // Set a lower memory limit for multipart forms (default is 32 MiB)
    // router.MaxMultipartMemory = 8 << 20  // 8 MiB
    router.POST("/upload", func(c *gin.Context) {
        // single file
        file, _ := c.FormFile("file")
        log.Println(file.Filename)

        // Upload the file to specific dst.
        c.SaveUploadedFile(file, file.Filename)

        /*
        也可以直接使用io操作，拷贝文件数据。
        out, err := os.Create(filename)
        defer out.Close()
        _, err = io.Copy(out, file)
        */

        c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
    })
    router.Run(":8080")
}
```



**2.2.5 Grouping routes**

router group是为了方便一部分相同的URL的管理，新建一个go文件(demo08_group.go)，

```go
v1 := router.Group("/v1")
    {
        v1.GET("/login", loginEndpoint)
        v1.GET("/submit", submitEndpoint)
        v1.POST("/read", readEndpoint)
    }
```



### Page 3  Model

#### 3.1 数据解析绑定

模型绑定可以将请求体绑定给一个类型。简单来说,，就是根据Body数据类型，将数据赋值到指定的结构体变量中 (类似于序列化和反序列化) 。

Gin提供了**两套绑定方法**：

- Must bind 
  - 方法：Bind`,`BindJSON`,`BindXML`,`BindQuery`,`BindYAML
- Should bind 
  - 方法：ShouldBind`,`ShouldBindJSON`,`ShouldBindXML`,`ShouldBindQuery`,`ShouldBindYAML
  - 行为：这些方法使用ShouldBindWith。如果存在绑定错误，则返回错误，开发人员有责任适当地处理请求和错误。

#### 3.2 JSON绑定   （这部分不是很懂）

JSON的绑定，其实就是将request中的Body中的数据按照JSON格式进行解析，解析后存储到结构体对象中。

```go
package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

type Login struct {
    User     string `form:"username" json:"user" uri:"user" xml:"user"  binding:"required"`
    Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}
//binding中的required 表示必须传递的参数

func main() {
    router := gin.Default()
    //1.binding JSON
    // Example for binding JSON ({"user": "hanru", "password": "hanru123"})
    router.POST("/loginJSON", func(c *gin.Context) {
        var json Login
        //其实就是将request中的Body中的数据按照JSON格式解析到json变量中
        if err := c.ShouldBindJSON(&json); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        //假如只传入一个User，那么会报错
        if json.User != "hanru" || json.Password != "hanru123" {
            c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
            return
        }
        c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
    })

    router.Run(":8080")
}
```

​		前面我们使用c.String返回响应，顾名思义则返回string类型。content-type是plain或者text。调用***c.JSON***   则返回json数据。其中  ***gin.H***  封装了生成json的方式，是一个强大的工具。使用golang可以像动态语言一样写字面量的json，对于嵌套json的实现，嵌套gin.H即可。



#### 3.3 Form表单

其实本质是将c中的request中的body数据解析到form中。

#### 3.4 Uri绑定

```go
router.GET("/:user/:password", func(c *gin.Context) {
        var login Login
        if err := c.ShouldBindUri(&login); err != nil {
            c.JSON(400, gin.H{"msg": err})
            return
        }
        c.JSON(200, gin.H{"username": login.User, "password": login.Password})
    })
```



### Page 4  响应

#### 4.1 JSON/XML/YAML

```go
 r.GET("/someJSON", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
    })

    r.GET("/moreJSON", func(c *gin.Context) {
        // You also can use a struct
        var msg struct {
            Name    string `json:"user"`
            Message string
            Number  int
        }
        msg.Name = "hanru"
        msg.Message = "hey"
        msg.Number = 123
        // 注意 msg.Name 变成了 "user" 字段
        // 以下方式都会输出 :   {"user": "hanru", "Message": "hey", "Number": 123}
        c.JSON(http.StatusOK, msg)
    })
//使用JSON，另外还有XML、YAML等
```

上面的代码功能是根据不同的请求，响应为不同的渲染方法（json，xml，yaml等）

#### 4.2 HTML模板渲染

gin支持加载HTML模板, 然后根据模板参数进行配置并返回相应的数据。

先要使用 LoadHTMLGlob() 或者 LoadHTMLFiles()方法来加载模板文件，

#### 4.3 文件响应

```go
router.StaticFS("/showDir", http.Dir("."))
```

这是访问当前项目目录下的内容

```go
router.StaticFS("/files", http.Dir("/bin"))
```

这是访问操作系统/bin下的内容

```go
router.StaticFile("/image", "./assets/miao.jpg")
```

显示选定的图片

http.Dir("/public")是利用本地tmp目录实现一个文件系统；
http.FileServer(http.Dir("/public"))返回一个Handler，其用来处理访问本地"/tmp"文件夹的HTTP请求；

#### 4.4 重定向

```go
r.GET("/redirect", func(c *gin.Context) {
        //支持内部和外部的重定向
        c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com/")
    })
```

进行访问，我们可以看到通过访问的路径，可以重定向到百度地址。

#### 4.5 同步异步



### Page 5  中间件

#### 5.1 全局中间件







### Page 6 数据库

基于database/sql的CURD操作

​		对于Gin本身，并没有对数据库的操作，本文实现的是，通过http访问程序，然后操作mysql数据库库。

