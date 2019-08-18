package main

import (
    "github.com/gin-gonic/gin"
    "log"
    "time"
)

type Person struct {
    Name string `form:"name"`
    Address string `form:"address"`
    Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
}

func main() {
    r := gin.Default()
    r.GET("/testing", startPage)
    _ = r.Run(":5000")
}

func startPage(c *gin.Context) {
    var person Person

    if c.ShouldBind(&person) == nil {
        log.Println(person.Name)
        log.Println(person.Address)
        log.Println(person.Birthday)
    }
    c.String(200, "success")
}
//GET "http://localhost:5000/testing?name=appleboy&address=xyz&birthday=1992-03-15"