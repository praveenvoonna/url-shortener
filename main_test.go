package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"URL-Shortener-Infra-Cloud/handlers"

	"github.com/gin-gonic/gin"
)

func TestShortenURL(t *testing.T) {
	router := gin.Default()
	router.POST("/shorten", handlers.ShortenURL)

	validURL := "https://www.example.com"
	payload, _ := json.Marshal(map[string]string{"url": validURL})
	reqValid, _ := http.NewRequest("POST", "/shorten", bytes.NewBuffer(payload))
	reqValid.Header.Set("Content-Type", "application/json")
	wValid := httptest.NewRecorder()
	router.ServeHTTP(wValid, reqValid)

	if wValid.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, wValid.Code)
	}

}

func TestRedirectToURL(t *testing.T) {
	router := gin.Default()
	router.GET("/:url", handlers.RedirectToURL)

	existingShortenedURL := "sampleShortenedURL"
	handlers.ShortenedUrlMap[existingShortenedURL] = "https://www.example.com"
	reqExisting := httptest.NewRequest("GET", "/"+existingShortenedURL, nil)
	wExisting := httptest.NewRecorder()
	router.ServeHTTP(wExisting, reqExisting)

	if wExisting.Code != http.StatusMovedPermanently {
		t.Errorf("Expected status code %d, but got %d", http.StatusMovedPermanently, wExisting.Code)
	}

	nonExistingShortenedURL := "nonExistingShortenedURL"
	reqNonExisting := httptest.NewRequest("GET", "/"+nonExistingShortenedURL, nil)
	wNonExisting := httptest.NewRecorder()
	router.ServeHTTP(wNonExisting, reqNonExisting)

	if wNonExisting.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d, but got %d", http.StatusNotFound, wNonExisting.Code)
	}

}

func TestGetShortenedURLs(t *testing.T) {
	router := gin.Default()
	router.GET("/shortenedurls", handlers.GetShortenedURLs)

	mockURLMap := map[string]string{
		"https://www.example.com":  "sampleShortenedURL1",
		"https://www.facebook.com": "sampleShortenedURL2",
	}

	req := httptest.NewRequest("GET", "/shortenedurls", nil)
	w := httptest.NewRecorder()
	handlers.UrlMap = mockURLMap
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}

}

func TestGetTopDomains(t *testing.T) {
	router := gin.Default()
	router.GET("/metrices", handlers.GetTopDomains)

	mockDomainCounts := map[string]int{
		"example.com":       5,
		"google.com":        3,
		"stackoverflow.com": 1,
		"github.com":        2,
	}

	handlers.DomainCounts = mockDomainCounts

	req := httptest.NewRequest("GET", "/metrices", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, w.Code)
	}

}
