package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	port := "3000"

	if p := os.Getenv("PORT"); p != "" {
		port = p
	}

	fmt.Println("starting server on port: " + port)

	http.ListenAndServe(":"+port, mux)
}
