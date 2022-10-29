package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
	"github.com/yisar/wcd/api"
)



func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	t, err := template.ParseFiles("sse.html")
	if err != nil {
		log.Fatal("WTF dude, error parsing your template.")

	}

	t.Execute(w, "friend")

	log.Println("Finished HTTP request at", r.URL.Path)
}

func main() {
	b := &Broker{
		make(map[chan string]bool),
		make(chan (chan string)),
		make(chan (chan string)),
		make(chan string),
	}

	b.Start()

	http.Handle("/events/", b)
	go func() {
		for i := 0; ; i++ {

			b.messages <- fmt.Sprintf("%d - the time is %v", i, time.Now())

			log.Printf("Sent message %d ", i)
			time.Sleep(5e9)

		}
	}()

	http.Handle("/", http.HandlerFunc(handler))
	http.ListenAndServe(":8000", nil)
}