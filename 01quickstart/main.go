package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {
    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })
    r.GET("/hello", getHello)
    _ = r.Run()
}

func getHello(c *gin.Context) {
    c.String(http.StatusOK, "Hello World!")
}


