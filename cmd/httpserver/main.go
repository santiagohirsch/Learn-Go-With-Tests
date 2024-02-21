package main

import (
	"Learn-Go-With-Tests/adapters/httpserver"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(httpserver.Handler)
	http.ListenAndServe(":8080", handler)
}
