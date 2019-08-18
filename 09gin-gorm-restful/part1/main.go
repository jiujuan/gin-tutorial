package main

import (
    "fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "math/rand"
    "time"
)

var db *gorm.DB

//初始化数据
func init() {
    var err error
    db, err = gorm.Open("mysql","root:root@/apidemo?charset=utf8&parseTime=True&loc=Local")
    if err != nil {
        panic(err)
    }
}

type Article struct {
    Title string
    Body string
    CreateAt time.Time
    UpdateAt time.Time
    ID uint `gorm:"primary_key"`
}

func main() {
    //migrate the schema
    db.AutoMigrate(&Article{})

    //插入新数据
    insertdata := Article{Title: "new title "+fmt.Sprintf("%d", rand.Int()), Body: "new body 10"}
    db.Create(&insertdata)

    //查询数据
    var article Article
    db.First(&article)
    fmt.Println(article.ID, article.Body, article.Title, article.CreateAt)

    var article2 Article
    db.First(&article2, 3) //查询id=3数据
    fmt.Println(article2.ID, article2.Body, article2.Title, article2.CreateAt)

    //find 查询
    var article3 Article
    db.Find(&article3) //查询所有数据
    //db.Find(&article3, 3)  //查询id=3 的数据

    //where 查询
    fmt.Println("3 where: ")
    var article4 Article
    db.Where("title=?", "au").First(&article4) //获取第一个匹配记录
    fmt.Println(article4.Title)

    // like 查询
    fmt.Println("4 like: ")
    var article5 Article
    db.Where("title Like ?", "%quas%").Find(&article5)
    fmt.Println("id:", article5.ID,"title:", article5.Title)

    //where struct and map
    fmt.Println("")
    fmt.Println("where struct & map")

    //struct
    var article6 Article
    db.Where(&Article{ID: 38, Title: "Enim"}).First(&article6)
    fmt.Println(article6.ID, article6.Title)
    //map
    var article7 Article
    db.Where(map[string]interface{}{"id": 38, "title": "Enim"}).Find(&article7)
    fmt.Println(article7.ID, article7.Title)

    //删除
    err := db.Delete(&Article{}, 40)
    if err.Error != nil {
        panic(err.Error)
    }

    //更新
    //更新第一条记录
    var article10 Article
    db.First(&article10)
    article10.Title = "hello2"
    article10.Body = "world2"
    db.Save(&article10)

    //按条件更新
    db.Model(&Article{}).Where("id = ?", 42).Update("title", "update title2")
}
