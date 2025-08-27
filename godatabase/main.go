package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string
	Email string
}

type Post struct {
	gorm.Model
	Name   string
	UserId int
}

func main() {
	dns := "root:@tcp(127.0.0.1:3306)/tutorial_go"

	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&User{}, &Post{})

	if err != nil {
		panic(err)
	}
	//SELECT

	// user := User{Name: "Jordan", Email: "OXALC@GMAIL.com"}

	// db.Create(&user)
	fmt.Println("final")
}
