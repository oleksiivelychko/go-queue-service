package main

import (
	"github.com/oleksiivelychko/go-queue-service/initmq"
	"io"
	"log"
	"net/http"
)

func main() {
	initmq.LoadEnv()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = io.WriteString(w, "It's Skaffold!\n")
	})
	http.HandleFunc("/queue/", func(w http.ResponseWriter, r *http.Request) {
		SendMessageIntoQueue(w, r)
	})

	log.Println("starting server...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
