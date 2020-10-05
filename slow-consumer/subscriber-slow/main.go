package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"predic8.de/util"
)

func main() {

	p := message.NewPrinter(language.German)

	nc, _ := nats.Connect(nats.DefaultURL,nats.ErrorHandler(func(c *nats.Conn, s *nats.Subscription, err error) {
		fmt.Println(err)
	}))
	defer nc.Close()

	sub, _ := nc.Subscribe("updates", func(msg *nats.Msg) {
		queued, _ , _ := msg.Sub.Pending();
		dropped, _ := msg.Sub.Dropped();
		p.Printf("%s     Queued: %d,  Verworfen: %d\n", msg.Data, queued, dropped)
	})

	fmt.Printf("Subscribed %s\n", sub.Queue)

	util.Block()
}
