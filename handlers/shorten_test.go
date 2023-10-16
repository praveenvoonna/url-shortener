package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestShortenURL(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("POST", "/shorten", strings.NewReader(`{"url": "https://www.example.com"}`))

	ShortenURL(c)

	if w.Code != http.StatusOK && w.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d or %d but got %d", http.StatusOK, http.StatusBadRequest, w.Code)
	}

}
