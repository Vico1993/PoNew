package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Vico1993/Ponew/domain/handler"
	"github.com/Vico1993/Ponew/domain/notification"
	"github.com/gorilla/mux"
)

func main() {
	// Init our notification system
	notification.InitNotification()

	router := mux.NewRouter()
	router.HandleFunc("/notifications/{type}", handler.PostNotification).Methods("POST")

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Hello World")
	}).Methods("GET")

	// Check header middleware
	// router.Use(middleware.CheckRequest)

	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
