# gin+vue+gorm 简单应用

gin 结合 vue，gorm 的使用，写一个简单的 demo。

## 安装 gin

安装 gin：

```shell
go get github.com/gin-gonic/gin
```

## 编写 gin demo

### demo 开始

1.新建 router，handler 文件夹:

```shell
15vue-gorm
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
15vue-gorm
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

## 安装 gorm

### 安装 gorm

```go
go get github.com/go-sql-driver/mysql
go get github.com/jinzhu/gorm
```

### gorm demo 练习

gorm 做个简单的练习，在文件夹下新增一个 db 文件夹，在里面新建 gorm.go ，如下图：

```shell
15vue-gorm
|-- db
   |-- gorm.go
|-- router
   |- router.go
|-- handler
   |- hello.go
|-- templates
   |- index.html
main.go
```

在 gorm.go 里面写上代码：

```go
package db

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type dbengine struct {
	*gorm.DB
}

func newDB() *dbengine {
    // 写上连接 mysql 的用户名，密码，你新建的数据库名 test，最后是数据库编码
	dsn := fmt.Sprintf("%s:%s@/%s?charset=%s", "root", "root", "test", "utf8")
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Printf("open db error: %s", err.Error())
	}

	if err = db.DB().Ping(); err != nil {
		log.Printf("ping error: %s", err.Error())
	}

	db.Set("gorm:table_options", "ENGINE=InnoDB")
	db.LogMode(true)

	engine := new(dbengine)
	engine.DB = db
	return engine
}

func GetDB() *dbengine {
	return newDB()
}
```

然后在你的 mysql 数据库里新建数据库 test，这是我们下面会使用到的一个库。



在 handler 目录下新建 student.go ，代码如下：

```go
package handler

import (
	"gin-tutorial/15vue-gorm/db"
	"log"
)

type Student struct {
	ID     int    `gorm:"primary_key;not null"`
	Name   string `gorm:"type:varchar(30);not null`
	Info   string `gorm:type:varchar(300)`
	Status string `gorm:type:char(2);not null`
}

func CreateStudentTable() {
	gorm := db.GetDB()

	gorm.DB.AutoMigrate(&Student{})

	log.Println("create table: student")
}

func InsertStudent() {
	students := Student{
		Name:   "tom",
		Info:   "a top student",
		Status: "1",
	}

	gorm := db.GetDB()
	gorm.Create(&students)
}

func GetStudents() []Student {
	var students []Student

	gorm := db.GetDB()
	gorm.Order("ID DESC").Find(&students)
	return students
}
```

然后在 router.go 里面增加以下代码：

```go
v1.GET("/db/student/createtable", func(c *gin.Context) {
    handler.CreateStudentTable()
    c.String(http.StatusOK, "OK")
})

v1.GET("/db/student/insert", func(c *gin.Context) {
    handler.InsertStudent()
    c.String(http.StatusOK, "insert daat success!")
})
```

1.在浏览器上先运行： http://localhost:8080/v1/db/student/createtable，新建 students 数据库

2.然后在运行新：http://localhost:8080/v1/db/student/insert ， 插入数据

> 新建数据库不会这样在浏览器上运行，这里只是测试编写 gorm 的操作代码。

