package main

import (
	"embed"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//go:embed assets/*
	var f embed.FS

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website",
		})
	})

	router.GET("favicon.ico", func(c *gin.Context) {
		file, _ := f.ReadFile("assets/favicon.ico")
		c.Data(
			http.StatusOK,
			"image/x-icon",
			file,
		)
	})

	router.Run(":8080")
}
