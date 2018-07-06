package main

import (
	"log"
	"net/http"
)

func helloServer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is an example server.\n"))
}

func main() {
	log.Println("Server listen in 443 port. Please open https://localhost/hello")
	http.HandleFunc("/hello", helloServer)
	err := http.ListenAndServeTLS(":443", "ssl/localhost.pem", "ssl/localhost-key.pem", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
