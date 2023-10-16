package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestRedirectToURL(t *testing.T) {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Params = []gin.Param{{Key: "url", Value: "sampleShortenedURL"}}

	RedirectToURL(c)

	if w.Code != http.StatusOK && w.Code != http.StatusNotFound {
		t.Errorf("Expected status code %d or %d but got %d", http.StatusOK, http.StatusNotFound, w.Code)
	}

}
