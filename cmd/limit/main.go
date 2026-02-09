package main

import (
	"log"
	"net/http"
	"strconv"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Building response controller
	rc := http.NewResponseController(w)

	// Getting content length
	contentLength := r.Header.Get("Content-Length")
	length, _ := strconv.ParseInt(contentLength, 10, 64)

	// Validating if request body is present
	if length == 0 {
		_ = rc.SetReadDeadline(time.UnixMilli(0))

		w.WriteHeader(http.StatusUnprocessableEntity)
		_, _ = w.Write([]byte("Missing request body"))
		return
	}

	// Validating if request body is exceeding 1KB
	const bodyLimit = 1 << (10)
	if length > bodyLimit {
		_ = rc.SetReadDeadline(time.UnixMilli(0))

		w.WriteHeader(http.StatusRequestEntityTooLarge)
		_, _ = w.Write([]byte("Request body too large"))
		return
	}

	// Set deadline for the connection
	const timeout = 1 * time.Second
	err := rc.SetReadDeadline(time.Now().Add(timeout))
	if err != nil {
		return
	}

	_, _ = w.Write([]byte("Hello World!")) // Writing is working
}

func main() {
	server := http.Server{
		Addr:    ":8080",
		Handler: http.HandlerFunc(handler),
	}

	log.Println("HTTP server in running on :8080")
	log.Fatal(server.ListenAndServe())
}
