## 安装

先安装相关库：

```shell
$ go get -u github.com/swaggo/swag/cmd/swag
$ go get -u github.com/swaggo/gin-swagger
$ go get -u github.com/swaggo/files
$ go get -u github.com/alecthomas/template
```

查安装是否成功

```shell
swag -v
swag.exe version v1.7.0
```

（我的是 windows）

## 编写例子

### gin 的例子

main.go

```go
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.GET("/hello", Hello)
	_ = r.Run(":8080")
}

func Hello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "hello world!",
	})
}
```

运行 main.go

```shell
go run mian.go
```

然后在浏览器上运行：`http://localhost:8080/hello` ，显示内容如下：

```json
{"message":"hello world!","status":"OK"}
```

### gin swagger 例子

[gin-swagger](https://github.com/swaggo/gin-swagger) 的 README 上一个例子在加一点代码。在 14swagger 目录下新建一个 docs 的目录，

main.go

```go
package main

import (
	_ "gin-tutorial/14swagger/docs"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title gin-swagger test
// @version 0.1
// @description gin-swagger 测试内容
// @termsOfService test

// @license.name APACHE 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @contact.name 九卷
// @contact.url https://github.com/jiujuan
// @contact.email jiujuanfeng@163.com

// @host localhost:8080
// @BasePath /
func main() {
	r := gin.New()

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r.GET("/hello/:name", Hello)
	_ = r.Run(":8080")
}

// @Summary hell API
// @description hello接口
// @Produce json
// @Version 0.1
// @Accept json
// @Param name path string true "name"
// @Success 200 {object} string {"status":"OK","message":"hello world!"}
// @Router /hello/{name} [GET]
func Hello(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "hello world!",
		"data":    name,
	})
}

func getUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"message": "hello world!",
	})
}
```

1.进入 14swagger 目录，运行命令：`swag init`， 生成 api 文档相关文件

```shell
$ swag init
00:51:26 Generate swagger docs....
00:51:26 Generate general API Info, search dir:./
00:51:26 create docs.go at docs\docs.go
00:51:26 create swagger.json at docs\swagger.json
00:51:26 create swagger.yaml at docs\swagger.yaml
```

2.然后运行 `go run main.go`

3.在浏览器上执行 `http://localhost:8080/swagger/index.html`

就可以看到 swag 生成的 api 文档了.

