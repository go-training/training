package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

// HelloWorld for hello world
func HelloWorld() string {
	return "Hello World, golang workshop!"
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("handler http request")
	fmt.Fprintf(w, "%s, I love %s!", HelloWorld(), r.URL.Path[1:])
}

// handles pinging the endpoint and returns an error if the
// agent is in an unhealthy state.
func pinger(port string) error {
	resp, err := http.Get("http://localhost:" + port)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return fmt.Errorf("server returned non-200 status code")
	}
	return nil
}

func main() {
	var showVersion bool
	var ping bool
	var port string
	flag.BoolVar(&showVersion, "version", false, "Print version information.")
	flag.BoolVar(&showVersion, "v", false, "Print version information.")
	flag.BoolVar(&ping, "ping", false, "check server live.")
	flag.StringVar(&port, "port", "8080", "server port")
	flag.StringVar(&port, "p", "8080", "server port")
	flag.Parse()

	// Show version and exit
	if showVersion {
		fmt.Println("version: 1.0.0")
		os.Exit(0)
	}

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
