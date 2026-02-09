package main

import (
	"io"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	name, _ := io.ReadAll(r.Body)
	_, _ = w.Write([]byte("Hello! " + string(name)))
}

func main() {
	log.Println("HTTP server in running on :8080")
	log.Fatal(http.ListenAndServe(":8080", http.HandlerFunc(handler)))
}
