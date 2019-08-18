package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    r := gin.Default()

    r.GET("param/:name", getName)
    _ = r.Run()
}

func getName(c *gin.Context) {
    name := c.Param("name")
    c.String(http.StatusOK, "hello %s", name)
}