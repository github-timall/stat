package main

import (
	"net/http"
	"time"
	"log"
)

func checkErr(err error) {
	if err != nil {
		log.Fatalf("FAILURE: %v", err)
	}
}

func main() {
	r := NewRouter()

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}