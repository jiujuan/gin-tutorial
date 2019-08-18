package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "net/http"
    "time"
)

var db *gorm.DB

func init() {
    var err error
    db, err = gorm.Open("mysql","root:root@/apidemo?charset=utf8&parseTime=True&loc=Local")

    if err != nil {
        panic(err)
    }

    //migrate the schema
    db.AutoMigrate(&Article{})
    //db.LogMode(true)
}

type Article struct {
    Title string `json:"title"`
    Body string `json:"body"`
    CreateAt time.Time `json:"create_at"`
    UpdateAt time.Time `json:"update_at"`
    ID uint `gorm:"primary_key" json:"id"`
}

func main() {
    r := gin.Default()
    r.GET("/article/", getArticles)
    r.GET("/article/:id", getArticleOne)
    r.POST("/article", createArticle)
    r.PUT("/article/:id", updateArticle)
    r.DELETE("/article/:id", deleteArticle)

    _ = r.Run(":5000")
}

func getArticles(c *gin.Context) {
    articles := make([]Article, 10)
    db.Where("id > ?", 30).Find(&articles)
    c.JSON(http.StatusOK, gin.H{
        "data": articles,
        "code": 200,
        "msg": "success",
    })
}

func getArticleOne(c *gin.Context) {
    id := c.Param("id")
    var article Article
    if err := db.Where("id=?", id).Find(&article).Error; err != nil {
        c.AbortWithStatus(http.StatusNotFound)
        //c.JSON(http.StatusNotFound, gin.H{
        //    "error": "id not found",
        //})
        fmt.Println("id not found")
    } else {
        c.JSON(http.StatusOK, gin.H{
            "data": article,
            "code": 200,
            "msg": "success",
        })
    }
}

func createArticle(c *gin.Context) {
    var article Article
    if err := c.BindJSON(&article); err != nil {
        c.AbortWithStatus(http.StatusExpectationFailed)
        fmt.Print(err.Error())
    }
    article.CreateAt = time.Now()
    article.UpdateAt = time.Now()
    db.Create(&article)
    c.JSON(http.StatusOK, gin.H{
        "data": article,
        "code": 201,
        "msg": "success",
    })
}

func updateArticle(c *gin.Context) {
    id := c.Param("id")
    var article Article

    if err := db.Where("id =? ", id).First(&article).Error; err != nil {
        c.AbortWithStatus(http.StatusNotFound)
        fmt.Println(err)
    } else {
        if err := c.BindJSON(&article); err != nil {
            c.AbortWithStatus(http.StatusExpectationFailed)
            fmt.Print(err.Error())
        }
        article.UpdateAt = time.Now()
        db.Save(&article)
        c.JSON(http.StatusOK, gin.H{
            "data": article,
            "code": http.StatusOK,
            "msg": "success",
        })
    }
}

func deleteArticle(c *gin.Context) {
    id := c.Param("id")
    var article Article

    db.Where("id = ?", id).Delete(&article)
    c.JSON(http.StatusOK, gin.H{
        "code": 200,
        "msg": "success",
    })
}


