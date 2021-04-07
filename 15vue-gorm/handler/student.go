package handler

import (
	"gin-tutorial/15vue-gorm/db"
	"log"
)

const (
	NoStatus  = "0"
	YesStatus = "1"
)

type Student struct {
	ID     int    `gorm:"primary_key;not null" json:"id"`
	Name   string `gorm:"type:varchar(30);not null" json:"name"`
	Info   string `gorm:"type:varchar(300)" json:"info"`
	Status string `gorm:"type:char(2);not null" json:"status"`
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

func GetAllStudents() []Student {
	var students []Student

	gorm := db.GetDB()
	gorm.Order("ID DESC").Find(&students)
	return students
}

func GetStudent(id int) Student {
	var student Student

	gorm := db.GetDB()
	gorm.First(&student, id)
	return student
}

func ChangeStudentStatus(id int, status string) {
	var student Student

	gorm := db.GetDB()
	gorm.Model(&student).Where("id=?", id).Update("status", status)
}

func DeleteStudent(id int) {
	var student Student

	gorm := db.GetDB()
	gorm.Delete(&student, id)
}

func AddStudent(student *Student) {
	gorm := db.GetDB()
	gorm.Create(&student)
}
