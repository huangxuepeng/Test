package main

import (
	"context"
	"fmt"
	"pb/pb"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	name := "world"
	addr := "127.0.0.1:50051"
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("connect error to: ", addr)
	}
	defer conn.Close()

	// 实例化一个client对象传入参数
	c := pb.NewProductInfoClient(conn)

	// 初始化上下文，设置请求超时时间为一秒
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	//调用Sayhello方法
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		fmt.Println("can not get to : " + addr)
	} else {
		fmt.Println("response from server: ", r.GetMessage()+err.Error())
	}
}
