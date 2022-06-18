package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

//HandlerFunc
func indexHandler(c *gin.Context) {
	fmt.Println("index")
	c.JSON(http.StatusOK, gin.H{
		"msg": "index",
	})
}

//定义一个中间件
func m1(c *gin.Context) {
	fmt.Println("m1 in ...")
	// 计时
	start := time.Now()
	c.Next() //调用后续的处理函数

	//c.Abort() //阻止调用后续的处理函数
	cost := time.Since(start)
	fmt.Printf("cost:%v\n", cost)
	fmt.Printf("m1 out ...")
}

func m2(c *gin.Context) {
	fmt.Println("m2 in ...")
	c.Next() //调用后续的处理函数
	fmt.Println("m2 out ...")
}

func main() {
	r := gin.Default()

	// 点开Use 查看源码发现 Use(middleware ...HandlerFunc) IRoutes
	r.Use(m1, m2) //全局注册中间件函数m1 m2

	/*访问/index
	执行indexHandler之前 去执行注册的中间件  总的执行打印顺序是：m1 in  -> m2 in -> index -> m2 out -> m1 out
	*/
	r.GET("/index", indexHandler)
	r.GET("/shop", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "shop",
		})
	})
	r.GET("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "user",
		})
	})

	r.Run(":9090")
}
