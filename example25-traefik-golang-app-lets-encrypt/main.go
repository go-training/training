package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// HelloWorld for hello world
func HelloWorld() string {
	return "Hello World, traefik workshop!"
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Got http request. time: %v", time.Now())
	fmt.Fprintf(w, "I love %s!, %s", r.URL.Path[1:], HelloWorld())
}

func pinger(port string) error {
	resp, err := http.Get("http://localhost:" + port)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return fmt.Errorf("server returned not-200 status code")
	}

	return nil
}

func main() {
	var port string
	var ping bool
	flag.StringVar(&port, "port", "8080", "server port")
	flag.StringVar(&port, "p", "8080", "server port")
	flag.BoolVar(&ping, "ping", false, "check server live")
	flag.Parse()

	if p, ok := os.LookupEnv("PORT"); ok {
		port = p
	}

	if ping {
		if err := pinger(port); err != nil {
			log.Printf("ping server error: %v\n", err)
		}

		return
	}

	http.HandleFunc("/", handler)
	log.Println("http server run on " + port + " port")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
