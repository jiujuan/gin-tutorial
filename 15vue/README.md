# gin+vue 简单应用

gin 结合 vue 的使用，写一个简单的 demo。

## 安装 gin

安装 gin：

```shell
go get github.com/gin-gonic/gin
```

## 编写 gin demo

### demo 开始

1.新建 router，handler 文件夹:

```shell
15vue
|-- router
   |- router.go
|-- handler
   |- hello.go
main.go
```

2.编写 router/router.go :

```go
package router

import (
	"gin-tutorial/15vue/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router() {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		v1.GET("/hello", handler.Hello)
		v1.GET("/hello/:name", func(c *gin.Context) {
			name := c.Param("name")
			c.String(http.StatusOK, "Hello, %s", name)
		})
	}

	r.Run(":8080")
}
```

3.编写 handler/hello.go :

```go
package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "Ok",
		"message": "hello world",
	})
}
```

4.编写 main.go

```go
package main

import (
	"gin-tutorial/15vue/router"
)

func main() {
	router.Router()
}
```

运行 go run main.go ，然后在浏览器上执行 http://localhost:8080/v1/hello ，http://localhost:8080/v1/hello/gin 都会输出相应的值。

### 增加 html 模板

新建文件夹 templates ，在里面新建 index.html 文件。此时文件夹结构如下：

```shell
15vue
|-- router
   |- router.go
|-- handler
   |- hello.go
|-- templates
   |- index.html
main.go
```

index.html 内容：

```html
<!DOCTYPE html>
<html>
<head>
    <title>gin+vue demo</title>
</head>
<body>
<h1>{{ .title }}</h1>
</body>
</html>
```

然后在 router.go 里面增加如下代码：

```go
r.LoadHTMLGlob("templates/*")

v1.GET("/index", func(c *gin.Context) {
    c.HTML(http.StatusOK, "index.html", gin.H{
        "title": "hello, world!",
    })
})
```

Router() 函数完整代码：

```go
func Router() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")
	v1 := r.Group("/v1")
	{
		v1.GET("/hello", handler.Hello)

		v1.GET("/hello/:name", func(c *gin.Context) {
			name := c.Param("name")
			c.String(http.StatusOK, "Hello, %s", name)
		})

		v1.GET("/index", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"title": "hello, world!",
			})
		})
	}

	r.Run(":8080")
}
```



重新运行 go run main.go , 然后在浏览器上运行：http://localhost:8080/v2/index ， 就会出现：hello, world!。一切运行正常。

## 安装 vue



