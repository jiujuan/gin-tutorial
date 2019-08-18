package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

// 简单说明，Login结构体，通过binding: 定义了"required" 就必须绑定，
// 就是请求时候，必须带上该参数，还分别进行了form、json、xml类型
type Login struct {
    User string `form:"user" json:"user" xml:"user" binding:"required"`
    Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

func main() {
    r := gin.Default()

    r.POST("/loginXML", func(c *gin.Context) {
        var xml Login
        if err := c.ShouldBindXML(&xml); err != nil{
            c.JSON(http.StatusBadRequest, gin.H{
                "error": "Login information is not complete",
            })
            return
        }

        if xml.User != "manu" || xml.Password != "123" {
            c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
        }

        c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
    })

    _ = r.Run(":5000")
}
//<?xml version="1.0" encoding="UTF-8"?>
//<root>
//<user>user</user>
//<password>123</user>
//</root>

