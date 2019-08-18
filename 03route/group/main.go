package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    v1 := r.Group("/v1")
    {
        v1.POST("/login", loginEndpoint)
        v1.POST("/read", readEndpoint)
    }

    v2 := r.Group("/v2")
    {
        v2.POST("/login", loginEndpoint)
        v2.POST("/read", readEndpoint)
    }
    _ = r.Run(":5000")
}

func loginEndpoint(c *gin.Context) {
    fmt.Println("group route: login")
}

func readEndpoint(c *gin.Context) {
    fmt.Println("group route: read")
}
//http://localhost:5000/v1/login
//http://localhost:5000/v1/read
//http://localhost:5000/v2/login
//http://localhost:5000/v2/read