package main

import (
	"net/http"

	"github.com/yisar/wcd/api"
)

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// fmt.Println("前置操作")
		w.Header().Add("Access-Control-Allow-Origin", "*")
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

	api.GlobalChan = api.Broker{
		Clients:        make(map[chan string]bool),
		NewClients:     make(chan (chan string)),
		DefunctClients: make(chan (chan string)),
		Messages:       make(chan string),
	}

	api.GlobalChan.Start()

	http.Handle("/events/", &api.GlobalChan)
	http.Handle("/login", http.HandlerFunc(api.Login))
	http.Handle("/gitlab-callback", http.HandlerFunc(api.Callback))
	http.Handle("/health", middleware(http.HandlerFunc(api.Health)))
	http.Handle("/prepare", middleware(http.HandlerFunc(api.Prepare)))

	http.ListenAndServe(":4000", nil)
}
