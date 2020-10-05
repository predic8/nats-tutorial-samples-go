package main

import (
	"github.com/nats-io/nats.go"
	"strconv"
)

func main() {
	nc, _ := nats.Connect(nats.DefaultURL)

	defer nc.Close()

	for i :=0 ; true; i++ {
		nc.Publish("updates", []byte("Msg " + strconv.Itoa(i)))
	}

}