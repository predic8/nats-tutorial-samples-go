package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"math"
	"predic8.de/util"
	"strconv"
	"time"
)

func prime(z int64) bool {
	time.Sleep(time.Second)
	var i int64
	for i = 2; i <= int64(math.Sqrt(float64(z))); i++ {
		if z % i == 0 {
			return false
		}
	}
	return true
}

func parseInt(msg *nats.Msg) int64 {
	i, _ := strconv.ParseInt(string(msg.Data), 10, 64)
	return i
}

func main() {
	nc, _ := nats.Connect(nats.DefaultURL)

	defer nc.Close()

	nc.Subscribe("services.primzahl", func(msg *nats.Msg) {
		fmt.Printf("Got: %s\n", msg.Data)
		if prime(parseInt(msg)) {
			msg.Respond([]byte("Prime"))
		} else {
			msg.Respond([]byte("No"))
		}
	})

	util.Block()
}


