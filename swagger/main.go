/**
 * Created by dongBao on 10/06/2019
 */

package main

import (
	"net/http"

	_ "goma/docs"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	ginSwagger "github.com/swaggo/gin-swagger/swaggerFiles"
)

func main() {
	r := gin.New()

	// 创建路由组
	v1 := r.Group("/api/v1")

	v1.GET("/record", record)

	v1.GET("/sayHello", sayHello)

	// 文档界面访问URL
	// http://127.0.0.1:8080/swagger/index.html
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}

// @summary 测试swagger
// @description 黄雪朋的学习swagger的过程
// @accept json
// @produce json
// @tags 黄雪朋的测试接口
// @success 200
// @param mobile body string true "下面填写手机号码"
// @param password body string true "填写密码"
// @Router /api/v1/record [get]
func record(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}

// 创建新用户 godoc
// @summary 创建新用户
// @description 管理员创建新用户账号，可以指定用户角色和是否为管理员账户
// @accept json, xml
// @produce json, xml
// @tags admin
// @param name body string true "用户名，建议使用姓名拼音"
// @param nick_name body string true "用户昵称，请使用真实姓名"
// @param phone body string true "手机号"
// @param email body string true "邮箱"
// @param password body string true "密码"
// @param role body string true "用户角色，student 学生或 teacher 老师"  Enums(student, teacher)
// @param is_admin body bool true "是否是管理员"
// @success 200
// @router /api/v1/admin/create-user [post]
func sayHello(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, name+",helloWorld")
}
