package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {
	r := mux.NewRouter()

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8080",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
