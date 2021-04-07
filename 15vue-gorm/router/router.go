package router

import (
	"gin-tutorial/15vue-gorm/handler"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Router() {
	r := gin.Default()

	r.Static("/templates", "./templates")
	// r.LoadHTMLGlob("templates/views/*")
	r.StaticFS("/vue", http.Dir("./templates/views"))

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

		v1.GET("/student/getAll", func(c *gin.Context) {
			c.JSON(200, handler.GetAllStudents())
		})

		v1.GET("/student/get", func(c *gin.Context) {
			strid := c.Query("id")
			id, _ := strconv.Atoi(strid)
			c.JSON(200, handler.GetStudent(id))
		})

		v1.POST("/student/add", func(c *gin.Context) {
			name := c.PostForm("name")
			info := c.PostForm("info")

			var student = handler.Student{
				Name:   name,
				Info:   info,
				Status: handler.NoStatus,
			}
			handler.AddStudent(&student)
			c.JSON(200, "OK")
		})

		v1.POST("/student/changestatus", func(c *gin.Context) {
			id := c.PostForm("id")
			status := c.PostForm("status")

			stuid, _ := strconv.Atoi(id)
			changestatus := handler.NoStatus

			if status == handler.NoStatus {
				changestatus = handler.YesStatus
			}

			handler.ChangeStudentStatus(stuid, changestatus)

			c.JSON(200, "OK")
		})

		v1.POST("/student/delete", func(c *gin.Context) {
			id, _ := strconv.Atoi(c.PostForm("id"))

			handler.DeleteStudent(id)

			c.JSON(200, "OK")
		})

	}

	r.Run(":8080")
}
