package main

import (
	"fmt"
	"net/http"
)

type MyHandler struct {
}

func (handler *MyHandler) ServeHTTP(write http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(write, "MyHandler: Hello World!")
}

func main() {
	handler := MyHandler{}

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: &handler,
	}
	server.ListenAndServe()
}
