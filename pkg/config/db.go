package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func ConnectDB() {
	const dsn string = "root:root@127.0.0.1:3306/books-management"
	d, err := gorm.Open(mysql.Open("root:root@tcp(127.0.0.1:3306)/books-management?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
