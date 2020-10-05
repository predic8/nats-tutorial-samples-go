package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"predic8.de/util"
)

func main() {
	fmt.Println("Shop Subscriber")

	nc, err := nats.Connect(nats.DefaultURL, nats.UserInfo("a", "geheim"),
		nats.ErrorHandler(func(cn *nats.Conn, sub *nats.Subscription, e error) {
			fmt.Printf("Error: %s\n", e)
		}),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()


	nc.Subscribe("updates", func(msg *nats.Msg) {
		fmt.Printf("%s\n", msg.Data)
	})

	util.Block()

}
