package initmq

import (
	"testing"
)

func TestMQConnection(t *testing.T) {
	LoadEnv()

	_, err := MQ()
	if err != nil {
		t.Errorf("unable to init RabbitMQ connection: %s", err)
	}
}
