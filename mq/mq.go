package mq

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"os"
)

func New() (conn *amqp.Connection, err error) {
	conn, err = amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s",
		os.Getenv("MQ_USER"),
		os.Getenv("MQ_PASS"),
		os.Getenv("MQ_HOST"),
		os.Getenv("MQ_PORT"),
	))
	return
}

func Queue(ch *amqp.Channel, name string) (amqp.Queue, error) {
	return ch.QueueDeclare(
		name,  // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
}

func FailOnError(err error) {
	if err != nil {
		log.Panic(err)
	}
}
