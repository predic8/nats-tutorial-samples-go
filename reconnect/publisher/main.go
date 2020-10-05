package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"strconv"
	"time"
)

func main() {
	nc, _ := nats.Connect(nats.DefaultURL,
		nats.DisconnectErrHandler(func(conn *nats.Conn, err error) {
			fmt.Println("Disconnect")
		}),
		nats.ReconnectHandler(func(conn *nats.Conn) {
			fmt.Println("Reconnect")
		}),
		nats.ReconnectBufSize(100),
		)

	defer nc.Close()

	for i :=0 ; true; i++ {
		time.Sleep(time.Second)
		err := nc.Publish("updates", []byte("Msg " + strconv.Itoa(i)))
		if err != nil {
			fmt.Printf("Error %s\n", err)
			continue
		}

		fmt.Printf("Nachricht %d gesendet.\n", i)

	}

}