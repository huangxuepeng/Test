package main

import (
	"testing"

	utils "github.com/Valiben/gin_unit_test"
	"github.com/gin-gonic/gin"
)

type OrdinaryResponse struct {
	Errno  string `json:"errno"`
	Errmsg string `json:"errmsg"`
}

func init() {
	// 需要写到init函数中, 启动gin框架
	router := gin.Default()
	router.GET("/", Test)
	utils.SetRouter(router)

}
func TestTest(t *testing.T) {
	resp := OrdinaryResponse{}
	// 调用函数发起http请求
	err := utils.TestHandlerUnMarshalResp("GET", "/", "form", "", &resp)
	if err != nil {
		t.Errorf("出现错误, %v", err)
		return
	}
	// fmt.Println(resp)
}
