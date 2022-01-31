package initmq

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

func Connection(host, user, pass, port string) (conn *amqp.Connection, err error) {
	url := fmt.Sprintf("amqp://%s:%s@%s:%s", user, pass, host, port)
	conn, err = amqp.Dial(url)
	return
}

func MakeQueue(ch *amqp.Channel, name string) (amqp.Queue, error) {
	return ch.QueueDeclare(
		name,  // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}
