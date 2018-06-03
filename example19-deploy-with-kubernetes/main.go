package main

import (
	"log"
	"net/http"
	"os"
	"strings"
)

var version = "master"

func showVersion(w http.ResponseWriter, r *http.Request) {
	log.Println(version)
	w.Write([]byte(version))
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello, drone got the message: " + message
	log.Println(message)
	w.Write([]byte(message))
}

func main() {
	// use PORT environment variable, or default to 8080
	port := "8080"
	if fromEnv := os.Getenv("PORT"); fromEnv != "" {
		port = fromEnv
	}
	http.HandleFunc("/version", showVersion)
	http.HandleFunc("/", sayHello)
	log.Println("Listen server on " + port + " port")
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
