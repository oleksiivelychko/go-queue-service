package initmq

import (
	"github.com/streadway/amqp"
	"os"
)

func LoadEnv(mqHost string) {
	if mqHost == "" {
		mqHost = "rabbitmq"
	}

	_ = os.Setenv("MQ_HOST", mqHost)
	_ = os.Setenv("MQ_PORT", "5672")
	_ = os.Setenv("MQ_USER", "rabbit")
	_ = os.Setenv("MQ_PASS", "secret")
}

func MQ() (*amqp.Connection, error) {
	return Connection(
		os.Getenv("MQ_HOST"),
		os.Getenv("MQ_USER"),
		os.Getenv("MQ_PASS"),
		os.Getenv("MQ_PORT"),
	)
}
