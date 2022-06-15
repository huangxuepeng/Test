package main

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type S struct {
	Sno    string `gorm:"column:SNO;type:char(2);"`
	Sname  string `gorm:"column:SNAME;type:char(6);"`
	Status string `gorm:"column:STATUS;type:char(2);"`
	City   string `gorm:"column:CITY;type:char(4);"`
}

func (S) TableName() string {
	return "s"
}
func init() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn))
	db.AutoMigrate(&S{})
}

func main() {
	// fmt.Println("000")
}
