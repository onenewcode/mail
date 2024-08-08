package mq

import (
	"os"

	"github.com/nats-io/nats.go"
)

var (
	Nc  *nats.Conn
	err error
)

func Init() {
	Nc, err = nats.Connect(os.Getenv("NATS_HOST"))
	if err != nil {
		panic(err)
	}
}
