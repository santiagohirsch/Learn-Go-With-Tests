package main

import (
	"log"
	"net"

	"Learn-Go-With-Tests/adapters/grpcserver"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	grpcserver.RegisterGreeterServer(s, &grpcserver.GreetServer{})

	if err := s.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
