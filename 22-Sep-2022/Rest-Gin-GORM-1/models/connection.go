package models

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase2(){
	dsn := "root:root@tcp(127.0.0.1:3306)/MYDB?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err!= nil {
		panic("failed to connect database : ")
	}

	db.AutoMigrate(User{})

	DB = db
}