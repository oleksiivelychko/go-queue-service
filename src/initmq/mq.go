package initmq

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

func Connection(host, user, pass, port string) (conn *amqp.Connection, err error) {
	url := fmt.Sprintf("amqp://%s:%s@%s:%s", user, pass, host, port)

	conn, err = amqp.Dial(url)
	if err != nil {
		log.Panicf("%s: %s", "Failed to connect to RabbitMQ", err)
	}

	return
}
