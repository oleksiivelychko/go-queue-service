package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	log.Println("starting server...")
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		_, _ = io.WriteString(w, "It's Skaffold!\n")
	})
	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		_, _ = io.WriteString(w, "Hello Skaffold!\n")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
