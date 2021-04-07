package db

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type dbengine struct {
	*gorm.DB
}

func newDB() *dbengine {
	dsn := fmt.Sprintf("%s:%s@/%s?charset=%s", "root", "root", "test", "utf8")
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Printf("open db error: %s", err.Error())
	}

	if err = db.DB().Ping(); err != nil {
		log.Printf("ping error: %s", err.Error())
	}

	db.Set("gorm:table_options", "ENGINE=InnoDB")
	db.LogMode(true)

	engine := new(dbengine)
	engine.DB = db
	return engine
}

func GetDB() *dbengine {
	return newDB()
}
