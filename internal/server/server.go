package server

import (
	"context"
	proto "github.com/isalikov/goservice/pkg/api/grpc"
)

type clients struct {}

type instance struct {
	proto.UnimplementedGoServiceServer
	clients *clients
}

func (i *instance) ServiceMethod(_ context.Context, _ *proto.ServiceMethodRequest) (*proto.ServiceMethodResponse, error) {
	return &proto.ServiceMethodResponse{Message: "..."}, nil
}
