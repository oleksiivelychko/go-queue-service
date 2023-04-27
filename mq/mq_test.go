package mq

import (
	"os"
	"testing"
)

func TestMQ(t *testing.T) {
	_ = os.Setenv("MQ_HOST", "go-queue-service.local")
	_ = os.Setenv("MQ_PORT", "5672")
	_ = os.Setenv("MQ_USER", "rabbit")
	_ = os.Setenv("MQ_PASS", "secret")

	conn, err := New()
	if err != nil {
		t.Fatal(err)
	}

	ch, err := conn.Channel()
	if err != nil {
		t.Fatal(err)
	}

	_, err = Queue(ch, "test")
	if err != nil {
		t.Fatal(err)
	}

	_, err = ch.QueueDelete("test", false, false, false)
	if err != nil {
		t.Fatal(err)
	}
}
