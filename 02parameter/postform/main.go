package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    r := gin.Default()
    r.POST("/name", getName)
    _ = r.Run()
}

//Multipart/Urlencoded Form
func getName(c *gin.Context) {
    message := c.PostForm("message")
    name := c.DefaultPostForm("name", "guest")

    c.JSON(http.StatusOK, gin.H{
        "status": 200,
        "message": message,
        "name": name,
    })
}
