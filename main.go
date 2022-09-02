package main

import (
	"fmt"
<<<<<<< HEAD
	"sort"
)

func numberOfPairs(nums []int) []int {
	arr := map[int]int{}
	for _, v := range nums {
		arr[v]++
	}
	res := 0
	for _, v := range arr {
		res += (v / 2)
	}
	return []int{res, len(nums) - 2*res}
}

func maximumSum(nums []int) int {
	sort.Ints(nums)
	arr := map[int][]int{}
	exist := false
	for k, v := range sumNum(nums) {
		c := arr[v]
		c = append(c, k)
		arr[v] = c
		if len(arr[v]) > 1 {
			exist = true
		}
	}
	if !exist {
		return -1
	}
	res := 0
	for _, v := range arr {
		if len(v) >= 2 {
			res = max(res, nums[v[len(v)-1]]+nums[v[len(v)-2]])
		}
	}
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func sumNum(nums []int) (arr []int) {
	// 将传来的数组进行转为数位的和, 顺序不改变
	for _, v := range nums {
		res := 0
		for v > 0 {
			res += v % 10
			v /= 10
		}
		arr = append(arr, res)
	}
	return arr
}
func main() {
	fmt.Println(maximumSum([]int{11, 3, 2, 1, 3, 2, 2}))

	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 实现两个数据库的关联, 并且查看数据关联的时候删除的情况
// user
type User struct {
	ID        int            `gorm:"primaryKey"`
	CreatedAt time.Time      `gorm:"column:create_time"`
	UpdatedAt time.Time      `gorm:"column:update_time"`
	DeletedAt gorm.DeletedAt `gorm:"<-:update"`
	Name      string         `gorm:"column:name;type:varchar(20);comment '姓名'"`
	Messages  []Message
}

// 信息表
type Message struct {
	Id        int            `gorm:"primaryKey"`
	CreatedAt time.Time      `gorm:"column:create_time"`
	UpdatedAt time.Time      `gorm:"column:update_time"`
	DeletedAt gorm.DeletedAt `gorm:"<-:update"`
	Data      string         `gorm:"column:data;type:varchar(30) "`
	UserID    int
}

var DB *gorm.DB

func init() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	//全局模式
	var err error
	DB, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		log.Println("连接数据库失败")
	}
	DB.AutoMigrate(
		&User{},
		&Message{},
	)
}
func main() {

	// user, msg := User{
	// 	Name: "哈哈哈",
	// },
	// 	Message{
	// 		Data:   "这是测试",
	// 		UserID: 1,
	// 	}
	// // 开始实现数据的插入
	// DB.Create(&user)
	// DB.Create(&msg)
	user := &User{
		ID: 1,
	}
	dd := DB.Select("Messages").Delete(&user)
	if dd.Error != nil {
		fmt.Println(dd.Error)
		return
	}
>>>>>>> df07758... 测试删除关联数据
}
