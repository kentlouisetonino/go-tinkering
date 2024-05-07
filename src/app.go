package main

import (
	"log"
	"net/http"
	"time"

	"github.com/kentlouisetonino/go-integration/src/routes"
)

func main() {
	http.HandleFunc("/api/health", routes.HealthRoute)

	s := &http.Server{
		Addr:           ":11000",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())
}
