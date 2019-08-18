package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.New()

    r.Use(gin.Logger())
    r.Use(gin.Recovery())

    //路由经过多个函数
    r.GET("/index", index(), chain)

    _ = r.Run(":5000")
}

func index() gin.HandlerFunc {
    return func(c *gin.Context) {
        fmt.Println("before middleware")
    }
}

func chain(c *gin.Context) {
    fmt.Println("index chainFunc")
}
