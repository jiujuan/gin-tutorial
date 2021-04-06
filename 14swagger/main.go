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
