package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := ":" + os.Getenv("PORT")
	stage := os.Getenv("UP_STAGE")

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong " + stage,
		})
	})

	r.GET("/v1", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong " + stage + " v1 ++ drone",
		})
	})

	r.GET("/v2", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong " + stage + " v1 ++ production",
		})
	})

	r.Run(port)
}
