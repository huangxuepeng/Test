// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: proto.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ProductInfoClient is the client API for ProductInfo service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductInfoClient interface {
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}

type productInfoClient struct {
	cc grpc.ClientConnInterface
}

func NewProductInfoClient(cc grpc.ClientConnInterface) ProductInfoClient {
	return &productInfoClient{cc}
}

func (c *productInfoClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error) {
	out := new(HelloReply)
	err := c.cc.Invoke(ctx, "/pb.ProductInfo/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductInfoServer is the server API for ProductInfo service.
// All implementations must embed UnimplementedProductInfoServer
// for forward compatibility
type ProductInfoServer interface {
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	mustEmbedUnimplementedProductInfoServer()
}

// UnimplementedProductInfoServer must be embedded to have forward compatible implementations.
type UnimplementedProductInfoServer struct {
}

func (UnimplementedProductInfoServer) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedProductInfoServer) mustEmbedUnimplementedProductInfoServer() {}

// UnsafeProductInfoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductInfoServer will
// result in compilation errors.
type UnsafeProductInfoServer interface {
	mustEmbedUnimplementedProductInfoServer()
}

func RegisterProductInfoServer(s grpc.ServiceRegistrar, srv ProductInfoServer) {
	s.RegisterService(&ProductInfo_ServiceDesc, srv)
}

func _ProductInfo_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductInfoServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ProductInfo/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductInfoServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductInfo_ServiceDesc is the grpc.ServiceDesc for ProductInfo service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductInfo_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ProductInfo",
	HandlerType: (*ProductInfoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _ProductInfo_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto.proto",
}
