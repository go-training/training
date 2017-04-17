package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ginEngine() *gin.Engine {
	router := gin.Default()

	router.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname")

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": gin.H{"data1": "data2"}})
	})

	return router
}

func main() {
	r := ginEngine()
	r.Run(":8080")
}
