package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/foobar", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "foobar",
		})
	})
	r.GET("/", func(c *gin.Context) {
		c.AbortWithStatus(http.StatusOK)
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
