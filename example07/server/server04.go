// +build 1.8

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, web 2")
}

func main() {
	server := &http.Server{
		Addr:         ":4000",
		WriteTimeout: 2 * time.Second,
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		if err := server.Close(); err != nil {
			log.Fatal("Server Closed:", err)
		}
	}()

	mux := http.NewServeMux()

	mux.Handle("/", &myHandle{})
	mux.HandleFunc("/hello", helloHandler)
	server.Handler = mux

	log.Println("Start server and listen on 4000")
	if err := server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			log.Println("Server closed under request")
		} else {
			log.Fatal(err)
		}
	}

	log.Println("Server exit")
}

type myHandle struct{}

func (*myHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome Server 2: "+r.URL.String())
}
