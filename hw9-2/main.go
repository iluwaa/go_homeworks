package main

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	zlog "github.com/rs/zerolog/log"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/register", register).Methods("POST")
	r.HandleFunc("/students", Authorized(students)).Methods("GET")
	r.HandleFunc("/students/{uid}", getTasks).Methods("GET")

	Class["0"] = &Participant{
		Name:    "Place",
		Surname: "Holder",
	}

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	zlog.Err(srv.ListenAndServe())
}
