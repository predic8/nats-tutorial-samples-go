package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"predic8.de/util"
)

func dump(conn *nats.Conn)  {
	fmt.Printf("Status: %s\n", conn.Status());
}

func main() {
	nc, err := nats.Connect(nats.DefaultURL,
			nats.DisconnectErrHandler(func(conn *nats.Conn, err error) {
				fmt.Printf("Disconnect")
				dump(conn)
			}),
			nats.ReconnectHandler(func(conn *nats.Conn) {
				fmt.Printf("Reconnect")
				dump(conn)
			}),
			nats.ErrorHandler(func(conn *nats.Conn, subscription *nats.Subscription, err error) {
				fmt.Printf("Error %s!", err)
				dump(conn)
			}),
			nats.ClosedHandler(func(conn *nats.Conn) {
				fmt.Println("Closed!")
				dump(conn);
			}),
			nats.DiscoveredServersHandler(func(conn *nats.Conn) {
				fmt.Println("Server entdeckt!")
				dump(conn);
			}),
		)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Servers: %s\n", nc.DiscoveredServers());

	fmt.Printf("# Subs: %s\n",nc.NumSubscriptions());


	defer nc.Close()

	nc.Subscribe("updates", func(msg *nats.Msg) {
		fmt.Printf("%s\n", msg.Data)
	})

	util.Block()
}
