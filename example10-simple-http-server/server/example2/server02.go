package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, web 2")
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", &myHandle{})
	mux.HandleFunc("/hello", helloHandler)

	log.Println("Start server and listen on 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}

type myHandle struct{}

func (*myHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome Server 2: "+r.URL.String())
}
