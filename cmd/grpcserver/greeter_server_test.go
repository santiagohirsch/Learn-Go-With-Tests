package main_test

import (
	"fmt"
	"testing"

	"Learn-Go-With-Tests/adapters"
	"Learn-Go-With-Tests/adapters/grpcserver"

	"Learn-Go-With-Tests/specifications"
)

func TestGreeterServer(t *testing.T) {
	var (
		port   = "50052"
		driver = grpcserver.Driver{Addr: fmt.Sprintf("localhost:%s", port)}
	)

	adapters.StartDockerServer(t, port, "grpcserver")
	specifications.GreetSpecification(t, &driver)
}
