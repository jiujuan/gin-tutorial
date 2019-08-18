package main

import (
    "github.com/gin-gonic/gin"
    "io"
    "os"
)

func main() {
    f, _ := os.Create("gin.log")
    gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        c.String(200, "pong")
    })

    _ = r.Run(":5000")
}
//可以用postman来测试