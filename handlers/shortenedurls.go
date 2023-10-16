package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetShortenedURLs(c *gin.Context) {
	c.JSON(http.StatusOK, UrlMap)
}
