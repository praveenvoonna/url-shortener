package main

import (
	"github.com/gin-gonic/gin"
	"github.com/praveenvoonna/URL-Shortener-Infra-Cloud/handlers"
)

func main() {
	r := gin.Default()

	// Routes
	r.POST("/shorten", handlers.ShortenURL)
	r.GET("/:url", handlers.RedirectToURL)
	r.GET("/shortenedurls", handlers.GetShortenedURLs)

	// Run the server
	r.Run(":8080")
}
