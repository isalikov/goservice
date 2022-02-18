package service

import (
	"context"
	"fmt"
	proto "github.com/isalikov/goservice/pkg/api/grpc"
	"google.golang.org/grpc"
	"time"
)

func Connect(addr string, t time.Duration) (proto.GoServiceClient, error) {
	timeout := time.Second * 5
	if t != 0 {
		timeout = t
	}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	conn, err := grpc.DialContext(
		ctx,
		addr,
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		return nil, &grpcError{
			Message: err.Error(),
			Service: fmt.Sprintf("Service: %v, URL: %v", "goservice", addr),
		}
	}

	return proto.NewGoServiceClient(conn), nil
}
