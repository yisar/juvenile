package main

import (
	"fmt"
	"github.com/yisar/wcd/api"
	"log"
	"net/http"
	"time"
)

func main() {
	b := api.Broker{
		Clients:        make(map[chan string]bool),
		NewClients:     make(chan (chan string)),
		DefunctClients: make(chan (chan string)),
		Messages:       make(chan string),
	}

	b.Start()

	http.Handle("/events/", &b)
	go func() {
		for i := 0; ; i++ {

			b.Messages <- fmt.Sprintf("%d - the time is %v", i, time.Now())

			log.Printf("Sent message %d ", i)
			time.Sleep(1e9)

		}
	}()

	http.ListenAndServe(":8000", nil)
}
