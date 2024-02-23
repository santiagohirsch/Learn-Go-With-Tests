package main

import (
	"Learn-Go-With-Tests/adapters/httpserver"
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", httpserver.NewHandler()))
}
