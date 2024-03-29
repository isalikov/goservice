// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: api/grpc/goservice.proto

package service

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

// GoServiceClient is the client API for GoService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GoServiceClient interface {
	ServiceMethod(ctx context.Context, in *ServiceMethodRequest, opts ...grpc.CallOption) (*ServiceMethodResponse, error)
}

type goServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGoServiceClient(cc grpc.ClientConnInterface) GoServiceClient {
	return &goServiceClient{cc}
}

func (c *goServiceClient) ServiceMethod(ctx context.Context, in *ServiceMethodRequest, opts ...grpc.CallOption) (*ServiceMethodResponse, error) {
	out := new(ServiceMethodResponse)
	err := c.cc.Invoke(ctx, "/goservice.GoService/ServiceMethod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GoServiceServer is the server API for GoService service.
// All implementations must embed UnimplementedGoServiceServer
// for forward compatibility
type GoServiceServer interface {
	ServiceMethod(context.Context, *ServiceMethodRequest) (*ServiceMethodResponse, error)
	mustEmbedUnimplementedGoServiceServer()
}

// UnimplementedGoServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGoServiceServer struct {
}

func (UnimplementedGoServiceServer) ServiceMethod(context.Context, *ServiceMethodRequest) (*ServiceMethodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ServiceMethod not implemented")
}
func (UnimplementedGoServiceServer) mustEmbedUnimplementedGoServiceServer() {}

// UnsafeGoServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GoServiceServer will
// result in compilation errors.
type UnsafeGoServiceServer interface {
	mustEmbedUnimplementedGoServiceServer()
}

func RegisterGoServiceServer(s grpc.ServiceRegistrar, srv GoServiceServer) {
	s.RegisterService(&GoService_ServiceDesc, srv)
}

func _GoService_ServiceMethod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServiceMethodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GoServiceServer).ServiceMethod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/goservice.GoService/ServiceMethod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GoServiceServer).ServiceMethod(ctx, req.(*ServiceMethodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GoService_ServiceDesc is the grpc.ServiceDesc for GoService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GoService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "goservice.GoService",
	HandlerType: (*GoServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ServiceMethod",
			Handler:    _GoService_ServiceMethod_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/grpc/goservice.proto",
}
