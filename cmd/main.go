package main

import (
	"fmt"
	"net/http"

	"github.com/M4nsur/snipURL/configs"
	"github.com/M4nsur/snipURL/internal/auth"
	"github.com/gorilla/mux"
)




func main() {

	conf := configs.LoadConfig()

	fmt.Println(conf.Db.DNS)


	fmt.Println("starting serv")
	router := mux.NewRouter()

	authHandler := auth.NewAuthHandlers()

	router.Path("/auth/login").Methods("POST").HandlerFunc(authHandler.Login)
	router.Path("/auth/register").Methods("POST").HandlerFunc(authHandler.Register)
	
	server := http.Server{
		Addr: ":9090",
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
		
}
