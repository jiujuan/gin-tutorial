package handler

import (
	"gin-tutorial/15vue-gorm/db"
	"log"
)

type Student struct {
	ID     int    `gorm:"primary_key;not null"`
	Name   string `gorm:"type:varchar(30);not null`
	Info   string `gorm:type:varchar(300)`
	Status string `gorm:type:char(2);not null`
}

func CreateStudentTable() {
	gorm := db.GetDB()

	gorm.DB.AutoMigrate(&Student{})

	log.Println("create table: student")
}

func InsertStudent() {
	students := Student{
		Name:   "tom",
		Info:   "a top student",
		Status: "1",
	}

	gorm := db.GetDB()
	gorm.Create(&students)
}

func GetStudents() []Student {
	var students []Student

	gorm := db.GetDB()
	gorm.Order("ID DESC").Find(&students)
	return students
}
