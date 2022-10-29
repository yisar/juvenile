package api

import(
	"net/http"
	"log"
	"fmt"
)

// sse 用于日志流传输，使用 channel 定时往 client 端传输消息

type Broker struct {
	clients        map[chan string]bool
	newClients     chan chan string
	defunctClients chan chan string
	messages       chan string
}

func (b *Broker) Start() {

	go func() {

		for {
			select {

			case s := <-b.newClients:
				b.clients[s] = true
				log.Println("Added new client")

			case s := <-b.defunctClients:
				delete(b.clients, s)
				close(s)
				log.Println("Removed client")

			case msg := <-b.messages:
				for s := range b.clients {
					s <- msg
				}
				log.Printf("Broadcast message to %d clients", len(b.clients))
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
	b.newClients <- messageChan

	notify := w.(http.CloseNotifier).CloseNotify()
	go func() {
		<-notify
		b.defunctClients <- messageChan
		log.Println("HTTP connection just closed.")
	}()

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Transfer-Encoding", "chunked")

	for {
		msg, open := <-messageChan

		if !open {
			break
		}
		fmt.Fprintf(w, "data: Message: %s\n\n", msg)

		f.Flush()
	}
	log.Println("Finished HTTP request at ", r.URL.Path)
}
