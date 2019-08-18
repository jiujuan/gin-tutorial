package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.New()

    //log
    r.Use(gin.Logger())
    //panic
    r.Use(gin.Recovery())

    //middleware
    authorized := r.Group("/author")
    authorized.Use(AuthRequired())
    {
        authorized.POST("/login", loginEndpoint)
    }

   _ = r.Run(":5000")
}
//curl -X POST http://localhost:5000/author/login
func loginEndpoint(c *gin.Context) {
    value, exist :=  c.Get("request")
    fmt.Println("authorized loginEndpoint --exist: ", exist, "--value: ", value)
}

func AuthRequired() gin.HandlerFunc {
    return func(c *gin.Context) {
        fmt.Println("middleware start")
        c.Set("request", "client_request")
        c.Next()
        fmt.Println("after middleware")
    }
}