package main

import (
	"embed"
	"github.com/oleksiivelychko/go-queue-service/initmq"
	"github.com/streadway/amqp"
	"html/template"
	"io"
	"log"
	"net/http"
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
	initmq.LoadEnv("rabbitmq")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = io.WriteString(w, "It's Skaffold!\n")
	})

	http.HandleFunc("/queue/", func(w http.ResponseWriter, r *http.Request) {
		page, ok := pages[r.URL.Path]
		if !ok {
			log.Printf("URL path %s not found", r.URL.Path)
			w.WriteHeader(http.StatusNotFound)
			return
		}

		tpl, err := template.ParseFS(embedTemplates, page)
		if err != nil {
			log.Printf("page %s not found", r.RequestURI)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)

		form := Form{
			Message: r.FormValue("message"),
			Sent:    false,
		}

		if err = tpl.Execute(w, form); err != nil {
			return
		}

		conn, err := initmq.MQ()
		initmq.FailOnError(err, "Failed to connect to RabbitMQ")
		defer conn.Close()

		ch, err := conn.Channel()
		initmq.FailOnError(err, "Failed to open a channel")
		defer ch.Close()

		q, err := initmq.MakeQueue(ch, "hello")
		initmq.FailOnError(err, "Failed to declare a queue")

		if r.Method == http.MethodPost {
			err = ch.Publish(
				"",     // exchange
				q.Name, // routing key
				false,  // mandatory
				false,  // immediate
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        []byte(form.Message),
				})

			initmq.FailOnError(err, "Failed to publish a message")
			log.Printf(" [x] Sent %s\n", form.Message)

			form.Sent = true
		}
	})

	http.FileServer(http.FS(embedTemplates))

	log.Println("starting server...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
