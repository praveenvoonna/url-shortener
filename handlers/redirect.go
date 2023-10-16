package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RedirectToURL(c *gin.Context) {
	shortened := c.Param("url")
	if original, exists := ShortenedUrlMap[shortened]; exists {
		c.Redirect(http.StatusMovedPermanently, original)
		return
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
}
