package main

import (
	"github.com/vavilen84/class_booking/handlers"
	"github.com/vavilen84/class_booking/store"
	"log"
)

func main() {
	store.InitDB()
	handler := handlers.MakeHandler()
	httpServer := handlers.InitHttpServer(handler)
	log.Fatal(httpServer.ListenAndServe())
}
