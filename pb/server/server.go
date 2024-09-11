package main

import (
	"context"
	"fmt"
	"net"
	"pb/pb"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedProductInfoServer
}

//	func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
//		return &pb.HelloReply{message: "hello" + in.name}, nil
//	}
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "hello" + in.Name}, nil
}
func main() {
	// 设置监听地址和端口
	lis, err := net.Listen("tcp", "127.0.0.1:50051")
	if err != nil {
		fmt.Println("failed to listen: ", err)
	}

	// 实例化1个服务器程序
	s := grpc.NewServer()

	// 调用服务注册函数
	pb.RegisterProductInfoServer(s, &server{})
	fmt.Println("server listening at", lis.Addr())

	// 在监听端口上运行服务器程序
	if err := s.Serve(lis); err != nil {
		fmt.Println("failed to serve: ", err)
	}
}
