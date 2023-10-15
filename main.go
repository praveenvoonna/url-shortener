package main

import (
	"URL-Shortener-Infra-Cloud/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Routes
	r.POST("/shorten", handlers.ShortenURL)
	r.GET("/:url", handlers.RedirectToURL)
	r.GET("/shortenedurls", handlers.GetShortenedURLs)
	r.GET("/metrices", handlers.GetTopDomains)

	// Run the server
	r.Run(":8080")
}
