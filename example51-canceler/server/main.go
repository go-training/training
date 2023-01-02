package main

import (
	"context"
	"net/http"
	"time"

	"example51/schedule"

	"github.com/gin-gonic/gin"
)

func main() {
	// initial schedule instance
	s := schedule.New()

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/cancel-task/:id", func(c *gin.Context) {
		taskID := c.Param("id")

		if err := s.Cancel(context.Background(), taskID); err != nil {
			c.String(http.StatusInternalServerError, "crash")
			return
		}

		c.String(http.StatusOK, "ok")
	})

	r.GET("/watch-task/:id", func(c *gin.Context) {
		taskID := c.Param("id")

		ctxDone, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		ok, _ := s.Cancelled(ctxDone, taskID)
		if ok {
			c.String(http.StatusOK, "true")
			return
		}

		c.String(http.StatusOK, "false")
	})

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
