package main

import (
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Tabler interface {
	TableName() string
}

// 图书
type Book struct {
	Number          string
	Category        string
	Publishinghouse string
	Author          string
	Title           string
	Pricing         string
}

type Reader struct {
	Numbering int
	Name      string
	Unit      string
	Sex       string
	Phone     string
}
type Borrow struct {
	Number    string
	Numbering int
	Time      string
}
type Student struct {
	Sno   string `gorm:"column:sno; comment '学号'"`
	Sname string `gorm:"column:name; comment '姓名'"`
	Ssex  string `gorm:"column:sex; comment '性别'"`
	Sage  int    `gorm:"column:age; comment '年龄'"`
	Sdept string `gorm:"column:dept; comment '所在系'"`
}

func (Student) TableName() string {
	return "Student"
}

type Course struct {
	Cno     int    `gorm:"column:cno; comment '课程名'"`
	Cname   string `gorm:"column:name; comment '课程名'"`
	Cpno    int    `gorm:"column:pno; comment '先行课'"`
	Ccredit int    `gorm:"column:credit; comment '学分'"`
}

func (Course) TableName() string {
	return "Course"
}

type SC struct {
	Sno   string `gorm:"column:sno; comment '学号'"`
	Cno   int    `gorm:"column:cno; comment '课程号'"`
	Grade int    `grom:"column:grade; comment '成绩'"`
}

func (SC) TableName() string {
	return "SC"
}
func (Book) TableName() string {
	return "book"
}
func (Reader) TableName() string {
	return "reader"
}
func (Borrow) TableName() string {
	return "borrow"
}
func init() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn))

	db.AutoMigrate(&Student{}, &Course{}, &SC{})
	ss := []Student{
		{Sno: "201215121", Sname: "李勇", Ssex: "男", Sage: 20, Sdept: "CS"},
		{Sno: "201215122", Sname: "刘晨", Ssex: "女", Sage: 19, Sdept: "CS"},
		{Sno: "201215123", Sname: "王敏", Ssex: "女", Sage: 18, Sdept: "MA"},
		{Sno: "201215125", Sname: "张立", Ssex: "男", Sage: 19, Sdept: "IS"},
	}
	dd := []Course{
		{Cno: 1, Cname: "数据库", Cpno: 5, Ccredit: 4},
		{Cno: 2, Cname: "数学", Cpno: 0, Ccredit: 2},
		{Cno: 3, Cname: "信息系统", Cpno: 1, Ccredit: 4},
		{Cno: 4, Cname: "操作系统", Cpno: 6, Ccredit: 3},
		{Cno: 5, Cname: "数据结构", Cpno: 7, Ccredit: 4},
		{Cno: 6, Cname: "数据处理", Cpno: 0, Ccredit: 2},
		{Cno: 7, Cname: "语言", Cpno: 6, Ccredit: 4},
	}
	cc := []SC{
		{Sno: "201215121", Cno: 1, Grade: 92},
		{Sno: "201215122", Cno: 2, Grade: 85},
		{Sno: "201215123", Cno: 3, Grade: 88},
		{Sno: "201215125", Cno: 2, Grade: 90},
		{Sno: "201215125", Cno: 3, Grade: 80},
	}
	db.Create(&ss)
	db.Create(&dd)
	db.Create(&cc)
}

func main() {
	// fmt.Println("000")
}
