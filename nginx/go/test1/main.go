package main

import "github.com/gin-gonic/gin"

/**
* 代码描述: 接口地址是 /
* 作者:小大白兔
* 创建时间:2022/11/15 10:35:05
 */
func main() {
	r := gin.Default()
	r.GET("test", test)
	r.Run(":8080")
}

func test(c *gin.Context) {
	c.JSON(200, gin.H{
		"code":    200,
		"message": "8080端口",
	})
}
