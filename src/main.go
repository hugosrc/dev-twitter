package main

import (
	"log"
	"net/http"
	"time"

	"github.com/hugosrc/dev-twitter/src/routes"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/users", routes.CreateUser).Methods("POST")

	server := &http.Server{
		Handler:      router,
		Addr:         "localhost:3333",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
