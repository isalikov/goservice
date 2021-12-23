package server

import (
	"fmt"
	"github.com/isalikov/goservice/internal/env"
	proto "github.com/isalikov/goservice/pkg/api/grpc"
	"google.golang.org/grpc"
	"net"
)

func Serve(cfg *env.Config) error {
	l, err := net.Listen("tcp", fmt.Sprintf(":%v", cfg.Port))
	if err != nil {
		return err
	}

	srv := &instance{ clients: &clients{}}

	grpcServer := grpc.NewServer()
	proto.RegisterGoServiceServer(grpcServer, srv)

	return grpcServer.Serve(l)
}
