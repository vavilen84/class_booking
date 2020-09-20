package handlers

import (
	"github.com/gorilla/mux"
	"github.com/vavilen84/class_booking/constants"
	"log"
	"net/http"
	"os"
)

func MakeHandler() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/api/class/{classId}", GetClass).Methods(http.MethodGet)
	//r.HandleFunc("/api/class", s.GetClassList).Methods(http.MethodGet)
	//r.HandleFunc("/api/class", s.CreateClass).Methods(http.MethodPost)
	//r.HandleFunc("/api/class/{classId}", s.UpdateClass).Methods(http.MethodPut)

	return r
}

func InitHttpServer(handler http.Handler) *http.Server {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	log.Printf("Listening on port %s", port)

	return &http.Server{
		Handler:      handler,
		Addr:         "127.0.0.1:" + port,
		WriteTimeout: constants.DefaultWriteTimout,
		ReadTimeout:  constants.DefaultReadTimeout,
	}
}
