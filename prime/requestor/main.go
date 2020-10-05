package main

import (
	"github.com/nats-io/nats.go"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"log"
	"strconv"
	"time"
)

func main() {

	p := message.NewPrinter(language.English)

	nc, _ := nats.Connect(nats.DefaultURL)
	defer nc.Close()

	for i :=100000000000000000 ; i < 1000000000000000000; i++ {
		res, err := nc.Request("services.primzahl", []byte(strconv.Itoa(i)) , 10 * time.Second);
		if err != nil {
			log.Print(err)
			continue
		}
		p.Printf("%d %s\n",i, res.Data)
	}

}