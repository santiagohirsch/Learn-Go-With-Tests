package grpcserver

import (
	"Learn-Go-With-Tests/domain/interactions"
	"context"
	"fmt"
)

type GreetServer struct {
	UnimplementedGreeterServer
}

func (g GreetServer) Greet(ctx context.Context, request *GreetRequest) (*GreetReply, error) {
	return &GreetReply{Message: interactions.Greet(request.Name)}, nil
}

func (g GreetServer) Curse(ctx context.Context, request *GreetRequest) (*GreetReply, error) {
	return &GreetReply{Message: fmt.Sprintf("Go to hell, %s!", request.Name)}, nil
}
