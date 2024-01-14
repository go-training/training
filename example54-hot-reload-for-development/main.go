package main

import (
	"log/slog"
	"net/http"
)

func main() {
	// write simple web server using http package
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	slog.Info("Server started at port 8080")
	// support grace shutdown
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
