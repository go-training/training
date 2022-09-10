package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Heartbeat for check server status
func Heartbeat(c *gin.Context) {
	c.AbortWithStatus(http.StatusOK)
}
