package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

/**
* 代码描述: 实现下载
* 作者:小大白兔
* 创建时间:2022/11/30 20:20:05
 */
func main() {
	r := gin.Default()
	r.GET("/test", Test)
	r.Run(":8080")
}

func Test(c *gin.Context) {
	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", "test.doc")) //fmt.Sprintf("attachment; filename=%s", filename)对下载的文件重命名
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	c.File("./项目.doc")
}
