package router

import (
	"gin-tutorial/15vue-gorm/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

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

		v1.GET("/db/student/createtable", func(c *gin.Context) {
			handler.CreateStudentTable()
			c.String(http.StatusOK, "OK")
		})

		v1.GET("/db/student/insert", func(c *gin.Context) {
			handler.InsertStudent()
			c.String(http.StatusOK, "insert daat success!")
		})
	}

	r.Run(":8080")
}
