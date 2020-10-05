package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"predic8.de/util"
)

func main() {

	nc, _ := nats.Connect(nats.DefaultURL)
	defer nc.Close()

	subject := "nrw.bonn"
	fmt.Printf("Subscribtion für %s\n", subject)

	nc.Subscribe(subject, func(m *nats.Msg) {
		fmt.Printf("%s\n", m.Data)
	})

	util.Block()
}
