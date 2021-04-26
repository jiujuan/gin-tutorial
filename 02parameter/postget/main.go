package main

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

//POST /post?id=1234&page=1 HTTP/1.1
//Content-Type: application/x-www-form-urlencoded
//name=manu&message=this_is_great
func main() {
    r := gin.Default()
    r.POST("/name", getName)
    _ = r.Run(":5000")
}

func getName(c *gin.Context) {
    id := c.Query("id")
    page := c.DefaultPostForm("page", "0")
    name := c.PostForm("name")
    message := c.PostForm("message")

    c.String(http.StatusOK, "id: %s, page: %s, name: %s, message: %s", id, page, name, message)
}
