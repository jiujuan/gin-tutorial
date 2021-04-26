package main

import (
	"fmt"

	"gin-tutorial/16hystrix-go/gin/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	fmt.Println("server start")
	router.Use(gin.Recovery())
	router.Use(middleware.HystrixDo)

	router.Handle("GET", "/hello", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "msg"})
	})

	router.Run(":8080")
}
