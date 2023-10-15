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
	for k, v := range domainCounts {
		pairs = append(pairs, Pair{k, v})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Value > pairs[j].Value
	})

	topDomains := make(map[string]int)
	for i := 0; i < 3 && i < len(pairs); i++ {
		topDomains[pairs[i].Key] = pairs[i].Value
	}

	c.JSON(http.StatusOK, topDomains)
}
