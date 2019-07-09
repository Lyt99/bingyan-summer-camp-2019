# gin notes

## Quick Start

**windows**

1. Install gin

   ```bash
   go get -u -v github.com/gin-gonic/gin
   ```

   this might be very slow...

   ```bash
   $env:GOPROXY = "https://goproxy.io"
   go get -u -v github.com/gin-gonic/gin
   ```

2. Install Govendor

   ```bash
   go get github.com/kardianos/govendor
   ```

3. Use Govendor

   ```bash
   (pwd)
   govendor init
   govendor fetch github.com/gin-gonic/gin@v1.3
   ```

4. *(Optimal) Copy a starting template inside project

   ```bash
   curl https://raw.githubusercontent.com/gin-gonic/examples/master/basic/main.go > main.go
   ```

   Another way to do so: create a file called `example.go`

   ```go
   package main
   
   import "github.com/gin-gonic/gin"
   
   func main() {
   	r := gin.Default()
   	r.GET("/ping", func(c *gin.Context) {
   		c.JSON(200, gin.H{
   			"message": "pong",
   		})
   	})
   	r.Run() // listen and serve on 0.0.0.0:8080
   }
   ```

## API Examples

### AsciiJSON

```go
func main() {
    router := gin.Default()
    
    router.GET("/someJSON", func (c *gin.Context) {
        data := map[string]interface{}{ // store any type of data in
            "lang": "go",				// interface{}
            "tag": "<br>",
        } // go style of json
        c.AsciiJSON(http.StatusOK, data)
    })
    
    router.run(":8080")
}
```

### Query and post form

For a post diagram:

```http
POST /post?id=1234&page=1 HTTP/1.1
Content-Type: application/x-www-form-urlencoded

name=manu&message=this_is_great
```

```go
func main() {
    router := gin.Default()
    
    router.POST("/post", func (c *gin.Context) {
        id := c.Query("id")
        page := c.DefaultQuery("page", "0")
        name := c.PostForm("name")
        message := c.PostForm("message")
        
        fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
    })
    router.Run(":8080")
}
```

```json
id: 1234; page: 1; name: manu; message: this_is_great
```

### Query string parameters

```go
func main() {
    router := gin.Default
    
    router.GET("/welcome", func (c *gin.Context) {
        firstname := c.DefaultQuery("firstname", "Guest")
        lastname := c.Query("lastname")
        
        c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
    })
    router.Run(":8080")
}
```

