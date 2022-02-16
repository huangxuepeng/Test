package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID   uint
	Name string `gorm:"type:varchar(10); not null; column:name"`
	Sex  int    `gorm:"column:sex;not null;default:1"`
	Card []Card
}
type Card struct {
	ID     uint
	Name   string `gorm:"type:varchar(10);not null;column:name"`
	haoma  string `gorm:"type:varchar(20);not null;column:hao_ma"`
	UserID uint
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
	if err := DB.AutoMigrate(&User{}, &Card{}); err != nil {
		log.Println("新建失败")
	}
	fmt.Println("成功")
}

func main() {
	var user User
	var card []Card
	user = User{
		Name: "jj",
		Sex:  1,
	}
	DB.Create(&User{Name: "jj", Card: []Card{{Name: "j", haoma: "345"}, {Name: "9", haoma: "567890"}}})
	DB.Model(&user).Association("Card").Find(&card)
	fmt.Println(user, card)
}
