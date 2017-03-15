package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, web 3")
}

func main() {
	server := &http.Server{
		Addr:         ":4000",
		WriteTimeout: 2 * time.Second,
	}

	mux := http.NewServeMux()

	mux.Handle("/", &myHandle{})
	mux.HandleFunc("/hello", helloHandler)
	server.Handler = mux

	log.Println("Start server and listen on 4000")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

type myHandle struct{}

func (*myHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome Server 3: "+r.URL.String())
}
