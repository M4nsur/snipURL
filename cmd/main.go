package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)


func testHttp(w http.ResponseWriter, r *http.Request) {
	fmt.Println("its work")
}


func main() {

	

	fmt.Println("starting serv")
	router := mux.NewRouter()

	router.Path("/").Methods("GET").HandlerFunc(testHttp)
	
	server := http.Server{
		Addr: ":9090",
		Handler: router,
	}


	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
		
}
