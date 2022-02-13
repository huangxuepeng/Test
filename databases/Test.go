package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
}

var (
	DB *gorm.DB
)

func init() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/huangxuepeng?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
		return
	}
	DB = db
	fmt.Println("成功")
}

// func main() {

// }
