package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"strconv"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL, nats.UserInfo("producer", "geheim"),
		nats.ErrorHandler(func(cn *nats.Conn, sub *nats.Subscription, e error) {
			fmt.Printf("Error: %s\n", e)
		}),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	for i := 0; i < 10; i++ {
		err := nc.Publish("updates", []byte("Msg "+strconv.Itoa(i)))
		if err != nil {
			fmt.Printf("Eroor")
			log.Fatal(err)
		}
	}
}
