package main

import (
	"fmt"
	"log"
	"net/http"
)

// Simple HTTP server: responds with a greeting on "/"
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, Tekton Web-test!")
}

func main() {
	http.HandleFunc("/", hello)
	addr := ":8080"
	log.Printf("starting server at %s", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
