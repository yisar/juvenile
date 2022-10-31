package api

import (
	"fmt"
	"log"
	"net/http"
)

// sse 用于日志流传输，使用 channel 定时往 client 端传输消息

type Broker struct {
	Clients        map[chan string]bool
	NewClients     chan chan string
	DefunctClients chan chan string
	Messages       chan string
}

var GlobalChan Broker

func (b *Broker) Start() {
	go func() {
		for {
			select {
			case s := <-b.NewClients:
				b.Clients[s] = true
				log.Println("Added new client")
			case s := <-b.DefunctClients:
				delete(b.Clients, s)
				close(s)
				log.Println("Removed client")
			case msg := <-b.Messages:
				for s := range b.Clients {
					s <- msg
				}
				log.Printf("Broadcast message to %d Clients", len(b.Clients))
			}
		}
	}()
}

func (b *Broker) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	f, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	messageChan := make(chan string)
	b.NewClients <- messageChan

	notify := w.(http.CloseNotifier).CloseNotify()
	go func() {
		<-notify
		b.DefunctClients <- messageChan
		log.Println("HTTP connection just closed.")
	}()

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Transfer-Encoding", "chunked")
	w.Header().Add("Access-Control-Allow-Origin", "*")

	for {
		msg, open := <-messageChan

		if !open {
			break
		}
		fmt.Fprintf(w, "data: > %s\n\n", msg)
		f.Flush()
	}
	log.Println("Finished HTTP request at ", r.URL.Path)
}
