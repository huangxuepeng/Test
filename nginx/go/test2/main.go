package main

import "github.com/gin-gonic/gin"

/**
* 代码描述: 接口的地址是 /api
* 作者:小大白兔
* 创建时间:2022/11/15 10:34:49
 */
func main() {
	r := gin.Default()
	r.GET("api/test", test)
	r.Run(":8081")
}

func test(c *gin.Context) {
	c.JSON(200, gin.H{
		"code":    200,
		"message": "8081端口",
	})
}

// docker run -itd --name nginx -p 8888:80 -p 443:443 -v docker/nginx/conf.d/nginx.conf:/etc/nginx/conf.d/nginx.conf -v docker/nginx/conf.d/nginx/cert:/etc/nginx -m 100m nginx