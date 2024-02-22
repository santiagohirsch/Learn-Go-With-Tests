package main_test

import (
	"Learn-Go-With-Tests/adapters"
	"Learn-Go-With-Tests/adapters/httpserver"
	"Learn-Go-With-Tests/specifications"
	"fmt"
	"net/http"
	"testing"
	"time"
)

func TestGreeterServer(t *testing.T) {
	var (
		port   = "8080"
		driver = httpserver.Driver{
			BaseURL: fmt.Sprintf("http://localhost:%s", port),
			Client: &http.Client{
				Timeout: 1 * time.Second,
			},
		}
	)

	adapters.StartDockerServer(t, port, "httpserver")
	specifications.GreetSpecification(t, &driver)
}
