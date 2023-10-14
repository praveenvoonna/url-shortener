package handlers

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

var urlMap = make(map[string]string)

func ShortenURL(c *gin.Context) {
	var json struct {
		URL string `json:"url"`
	}
	if err := c.ShouldBindJSON(&json); err == nil {
		if shortened, exists := urlMap[json.URL]; exists {
			c.JSON(http.StatusOK, gin.H{"shortened_url": shortened})
			return
		}
		shortened := generateRandomString(6)
		urlMap[shortened] = json.URL
		c.JSON(http.StatusOK, gin.H{"shortened_url": shortened})
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
