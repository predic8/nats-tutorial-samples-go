package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"strconv"
)

func main() {

	fmt.Println("Personal Publisher")

	nc, err := nats.Connect(nats.DefaultURL, nats.UserInfo("c", "geheim"),
		nats.ErrorHandler(func(cn *nats.Conn, sub *nats.Subscription, e error) {
			fmt.Printf("Error: %s\n", e)
		}),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	for i := 0; i < 10; i++ {
		err := nc.Publish("updates", []byte("Personal Updates "+strconv.Itoa(i)))
		if err != nil {
			fmt.Printf("Error")
			log.Fatal(err)
		}
	}
}
