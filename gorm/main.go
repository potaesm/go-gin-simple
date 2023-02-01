package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func connectDatabase() {
	database, err := gorm.Open(mysql.Open("sa:password@tcp(127.0.0.1:3306)/test?charset=utf8"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database!")
	}
	DB = database
}

func main() {
	connectDatabase()
	fmt.Println("Database connected!")
}
