package handlers

import (
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
)

func GetTopDomains(c *gin.Context) {
	type Pair struct {
		Key   string
		Value int
	}

	var pairs []Pair
	for k, v := range DomainCounts {
		pairs = append(pairs, Pair{k, v})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Value > pairs[j].Value
	})

	domSize := min(3, len(pairs))
	topDomains := make([]Pair, 0)
	for i := 0; i < domSize; i++ {
		topDomains = append(topDomains, pairs[i])
	}

	c.JSON(http.StatusOK, topDomains)
}
