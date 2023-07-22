package main

import (
	"context"
	"embed"
	"github.com/oleksiivelychko/go-queue-service/mq"
	amqp "github.com/rabbitmq/amqp091-go"
	"html/template"
	"io"
	"log"
	"net/http"
	"time"
)

var (
	//go:embed templates
	embedTemplates embed.FS
	pages          = map[string]string{
		"/queue/": "templates/send_message_form.html",
	}
)

type Form struct {
	Message string
	Sent    bool
}

func main() {
	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		_, _ = io.WriteString(resp, "It's Skaffold!\n")
	})

	http.HandleFunc("/queue/", func(resp http.ResponseWriter, req *http.Request) {
		page, ok := pages[req.URL.Path]
		if !ok {
			log.Printf("URL path %s not found", req.URL.Path)
			resp.WriteHeader(http.StatusNotFound)
			return
		}

		tpl, err := template.ParseFS(embedTemplates, page)
		if err != nil {
			log.Printf("page %s not found", req.RequestURI)
			resp.WriteHeader(http.StatusInternalServerError)
			return
		}

		resp.Header().Set("Content-Type", "text/html")
		resp.WriteHeader(http.StatusOK)

		form := Form{
			Message: req.FormValue("message"),
			Sent:    false,
		}

		if err = tpl.Execute(resp, form); err != nil {
			return
		}

		conn, err := mq.New()
		mq.FailOnError(err)
		defer func(conn *amqp.Connection) {
			_ = conn.Close()
		}(conn)

		ch, err := conn.Channel()
		mq.FailOnError(err)
		defer func(ch *amqp.Channel) {
			_ = ch.Close()
		}(ch)

		queue, err := mq.Queue(ch, "go-queue")
		mq.FailOnError(err)

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		if req.Method == http.MethodPost {
			err = ch.PublishWithContext(ctx, "", queue.Name, false, false,
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        []byte(form.Message),
				})

			mq.FailOnError(err)
			log.Printf(" [x] Sent: %s\n", form.Message)

			form.Sent = true
		}
	})

	http.FileServer(http.FS(embedTemplates))

	log.Println("starting server...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
