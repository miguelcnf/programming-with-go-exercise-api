package main

import (
	"log"
	"net/http"

	"api/internal/clock"
)

func main() {
	http.HandleFunc(clock.CurrentEndpoint, clock.CurrentHandler)

	err := http.ListenAndServe(":8000", nil)
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("error on http server: %v", err)
	}
}
