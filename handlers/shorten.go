package handlers

import (
	"math/rand"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

var urlMap, shortenedUrlMap, domainCounts = make(map[string]string), make(map[string]string), make(map[string]int)

func ShortenURL(c *gin.Context) {
	var json struct {
		URL string `json:"url"`
	}
	if err := c.ShouldBindJSON(&json); err == nil {
		if shortened, exists := urlMap[json.URL]; exists {
			c.JSON(http.StatusOK, gin.H{json.URL: shortened})
			return
		}
		shortened := generateRandomString(6)
		urlMap[json.URL] = shortened
		shortenedUrlMap[shortened] = json.URL

		// Extract domain
		u, err := url.Parse(json.URL)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL"})
			return
		}
		domain := u.Host

		// Update domain count
		domainCounts[domain]++

		c.JSON(http.StatusOK, gin.H{json.URL: shortened})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func generateRandomString(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
