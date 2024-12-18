package main

import (
	"log"
	"net/http"

	"github.com/luquor/TinyGopher/internal"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/url/", internal.ResolveHandler)

	port := ":8080"
	log.Printf("Starting server on %s...\n", port)
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatalf("Server failed to start: %v\n", err)
	}
}
