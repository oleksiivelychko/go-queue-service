package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		_, _ = io.WriteString(w, "Hello, Skaffold!")
	})
	log.Fatal(http.ListenAndServe(":80", nil))
}
