

跨域认证其中一个方案是服务器索性不保存 session 数据了，所有数据都保存在客户端，每次请求都发回服务器。JWT 就是这种方案的一个代表。

WT 的三个部分依次如下。

  Header.Payload.Signature，也就是

> - Header（头部）
> - Payload（负载）
> - Signature（签名）

***Header 部分***是一个 JSON 对象，描述 JWT 的元数据

***Payload 部分***也是一个 JSON 对象，用来存放实际需要传递的数据。

ps：注意，JWT 默认是不加密的，任何人都可以读到，所以不要把秘密信息放在这个部分。

***Signature 部分***是对前两部分的签名，防止数据篡改。

#### JWT的使用方法

客户端收到服务器返回的 JWT，可以储存在 Cookie 里面，也可以储存在 localStorage。

此后，客户端每次与服务器通信，都要带上这个 JWT。你可以把它放在 Cookie 里面自动发送，但是这样不能跨域，所以更好的做法是放在 HTTP 请求的头信息`Authorization`字段里面

另一种做法是，跨域的时候，JWT 就放在 POST 请求的数据体里面。



token、cookie、session的区别：

***cookie*** 是一个非常具体的东西，指的就是浏览器里面能永久存储的一种数据，仅仅是浏览器实现的一种数据存储功能。

服务器使用***session***把用户的信息临时保存在了服务器上，用户离开网站后session会被销毁。

***token***的意思是“令牌”，是用户身份的验证方式，最简单的token组成:uid(用户唯一的身份标识)、time(当前时间的时间戳)、sign(签名，由token的前几位+盐以哈希算法压缩成一定长的十六进制字符串，可以防止恶意第三方拼接token请求服务器)。

用户注册之后, 服务器生成一个 JWT token返回给浏览器, 浏览器向服务器请求数据时将 JWT token 发给服务器, 服务器用 signature 中定义的方式解码 JWT 获取用户信息.











```go
Authenticator: func(c *gin.Context) (interface{}, error)
//用来判断是否是正确用户，需要自己来定义
```

