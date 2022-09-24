package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func test(email string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("email", email)
	}
}
func test1(c *gin.Context) {
}
func test2(c *gin.Context) {
	dd, _ := c.Get("email")
	fmt.Println(dd)
}
func main() {
	r := gin.Default()
	r.GET("/test1", test("2695009886"), test1)
	r.GET("/test2", test2)
	r.Run(":8081")
}
