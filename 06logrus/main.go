package main

import (
    "gin-tutorial/06logrus/middle"
    "github.com/gin-gonic/gin"
    "github.com/sirupsen/logrus"
)

func main() {
    log := logrus.New()
    r := gin.New()
    r.Use(middle.Logger(log), gin.Recovery())

    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "status": 200,

            "message": "logger test msg!",
        })
    })

    _ = r.Run(":5000")
}
// cd ./06logrus
// go run main.go

