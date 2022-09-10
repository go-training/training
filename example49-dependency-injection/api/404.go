package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// NotFound represents the 404 page.
func NotFound(c *gin.Context) {
	c.JSON(
		http.StatusNotFound,
		gin.H{
			"code":  http.StatusNotFound,
			"error": "PAGE NOT FOUND",
		},
	)
}
