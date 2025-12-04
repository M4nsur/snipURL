package main

import (
	"fmt"
	"net/http"

	"github.com/M4nsur/snipURL/internal/auth"
)

func main() {
	// conf := config.LoadConfig()
	router := http.NewServeMux()
	auth.NewAuthHandler(router)
	server := http.Server{
		Addr: ":8081",
		Handler: router,
	}
	fmt.Println("server is listening on port 8081")
	server.ListenAndServe()
}
