package main

import (
	"log"
	"net/http"
	"strings"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message
	log.Println(message)
	w.Write([]byte(message))
}
func main() {
	http.HandleFunc("/", sayHello)
	log.Println("Listen server on 8080 port")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
