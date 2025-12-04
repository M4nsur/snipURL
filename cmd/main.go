package main

import (
	"fmt"
	"net/http"

	config "github.com/M4nsur/snipURL/configs"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Hello")
}

func main() {
	conf := config.LoadConfig()
	router := http.NewServeMux()
	router.HandleFunc("/hello", hello)
	server := http.Server{
		Addr: ":8081",
		Handler: router,
	}
	fmt.Println("server is listening on port 8081")
	server.ListenAndServe()
}
