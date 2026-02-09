package main

import (
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Hello World!"))
}

func main() {
	const readTimeout = 2 * time.Second

	server := http.Server{
		Addr:        ":8080",
		ReadTimeout: readTimeout,
		Handler:     http.HandlerFunc(handler),
	}

	log.Println("HTTP server in running on :8080")
	log.Fatal(server.ListenAndServe())
}
