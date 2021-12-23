package service

import "fmt"

type grpcError struct {
	Message	string
	Service	string
}

func (e *grpcError) Error() string {
	return fmt.Sprintf("%d - %s", e.Message, e.Service)
}
