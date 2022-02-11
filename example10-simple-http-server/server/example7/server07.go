package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/acme/autocert"
	"golang.org/x/sync/errgroup"
)

func main() {
	r := gin.Default()

	// Ping handler
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	var g errgroup.Group

	g.Go(func() error {
		return http.ListenAndServe(":http", http.RedirectHandler("https://example.com", 303))
	})
	g.Go(func() error {
		return http.Serve(autocert.NewListener("example.com"), r)
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
