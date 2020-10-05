package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"predic8.de/util"
)

func main() {

	nc, _ := nats.Connect("nats://localhost:1003",
		nats.DisconnectErrHandler(func(conn *nats.Conn, err error) {
			fmt.Printf("Disconnect\n")
		}),
		nats.ReconnectHandler(func(conn *nats.Conn) {
			fmt.Printf("Reconnect\n")
			fmt.Printf("Servers: %s\n", conn.DiscoveredServers());
			fmt.Printf("Server %s\n", conn.ConnectedUrl())

		}),
		nats.ErrorHandler(func(conn *nats.Conn, subscription *nats.Subscription, err error) {
			fmt.Printf("Error %s!", err)
		}),
		nats.ClosedHandler(func(conn *nats.Conn) {
			fmt.Println("Closed!")
		}),
		nats.DiscoveredServersHandler(func(conn *nats.Conn) {
			fmt.Println("Server entdeckt!")
			fmt.Printf("Servers: %s\n", conn.DiscoveredServers());
		}),
	)
	defer nc.Close()

	fmt.Printf("Servers: %s\n", nc.DiscoveredServers());
	fmt.Printf("Verbinde mich mit: %s\n", nc.ConnectedUrl())

	nc.Subscribe("updates", func(msg *nats.Msg) {
		fmt.Printf("%s\n", msg.Data)
	})

	util.Block()
}
