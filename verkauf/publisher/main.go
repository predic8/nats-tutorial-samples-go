package main

import (
	"github.com/nats-io/nats.go"
	"math/rand"
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

	for i :=0 ; i < 10; i++ {
		ec.Publish("verkauf.west.bonn", artikel{rand.Intn(9999),rand.Float32() * 99, rand.Intn(100)})
		ec.Publish("verkauf.west.köln", artikel{rand.Intn(9999),rand.Float32() * 99, rand.Intn(100)})
	}

}