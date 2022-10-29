package main

import (
	"net/http"

	"github.com/yisar/wcd/api"
)

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// fmt.Println("前置操作")
		next.ServeHTTP(w, r)
		// fmt.Println("后置操作")
	})
}

func root(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`hello`))
}

func main() {

	http.Handle("/", middleware(http.HandlerFunc(root)))

	b := api.Broker{
		Clients:        make(map[chan string]bool),
		NewClients:     make(chan (chan string)),
		DefunctClients: make(chan (chan string)),
		Messages:       make(chan string),
	}

	b.Start()

	http.Handle("/events/", &b)
	http.Handle("/login", http.HandlerFunc(api.Login))
	http.Handle("/gitlab-callback", http.HandlerFunc(api.Callback))

	// go func() {
	// 	for i := 0; ; i++ {

	// 		b.Messages <- fmt.Sprintf("%d - the time is %v", i, time.Now())

	// 		log.Printf("Sent message %d ", i)
	// 		time.Sleep(2e9)

	// 	}
	// }()
	http.ListenAndServe(":4000", nil)
}
