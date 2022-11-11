package main

import (
	"io"
	"net/http"
)

func handlePing(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "pong")
}
func handlePong(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "ping")
}

func main() {
	http.HandleFunc("/ping", handlePing)
	http.HandleFunc("/pong", handlePong)
	http.HandleFunc("/", handlePing)

	http.ListenAndServe(":8080", nil)
}
