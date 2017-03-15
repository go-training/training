package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hi, 一天學會 golang")

	// try changing the value of this url
	res, err := http.Get("https://golang.org")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res.Status)
}
