package initmq

import (
	"testing"
)

func TestMQConnection(t *testing.T) {
	LoadEnv("go-queue-service.local")

	_, err := MQ()
	if err != nil {
		t.Errorf("unable to init RabbitMQ connection: %s", err)
	}
}

func TestMakeQueue(t *testing.T) {
	LoadEnv("go-queue-service.local")

	conn, _ := MQ()
	ch, err := conn.Channel()
	if err != nil {
		t.Errorf("unable to create RabbitMQ channel: %s", err)
	}

	_, err = MakeQueue(ch, "test")
	if err != nil {
		t.Errorf("unable to make RabbitMQ queue: %s", err)
	}
}
