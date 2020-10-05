package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"predic8.de/util"
)

type artikel struct {
	Artikel int
	Price  float32
	Menge int
}

func main() {

	nc, _ := nats.Connect(nats.DefaultURL)
	defer nc.Close()

	ec, _ := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	defer ec.Close()

	ec.Subscribe("verkauf.>", func(s *artikel) {
		fmt.Printf("%d Stück vom Artikel %d zum Preis von %f\n", s.Menge, s.Artikel, s.Price)
	})

	util.Block()
}
