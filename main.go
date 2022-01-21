package main

import (
	"github.com/oleksiivelychko/go-queue-service/initmq"
	"github.com/streadway/amqp"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

type Form struct {
	Message string
	Success bool
}

func main() {
	initmq.LoadEnv()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = io.WriteString(w, "It's Skaffold!\n")
	})
	http.HandleFunc("/queue/", func(w http.ResponseWriter, r *http.Request) {
		wd, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
		tmpl := template.Must(template.ParseFiles(filepath.Join(wd, "./templates/send_message_form.html")))

		form := Form{
			Message: r.FormValue("message"),
		}

		if r.Method != http.MethodPost {
			form.Success = false
			_ = tmpl.Execute(w, form)
			return
		}

		conn, err := initmq.MQ()
		failOnError(err, "Failed to connect to RabbitMQ")
		defer conn.Close()

		ch, err := conn.Channel()
		failOnError(err, "Failed to open a channel")
		defer ch.Close()

		q, err := initmq.MakeQueue(ch, "hello")
		failOnError(err, "Failed to declare a queue")

		err = ch.Publish(
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(form.Message),
			})
		failOnError(err, "Failed to publish a message")
		log.Printf(" [x] Sent %s\n", form.Message)

		form.Success = false
		_ = tmpl.Execute(w, form)
	})

	log.Println("starting server...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
