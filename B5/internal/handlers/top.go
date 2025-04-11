package handlers

import (
	"B5/internal/store"

	"github.com/gin-gonic/gin"
)

type TopHandler struct {
	store store.RedisStore
}

func NewTopHandler(store store.RedisStore) *TopHandler {
	return &TopHandler{store: store}
}

func (t *TopHandler) Handle(c *gin.Context) {
	topUsers, err := t.store.GetTopUsers(c, 10)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get top users"})
		return
	}
	c.JSON(200, gin.H{
		"top_users": topUsers,
	})
}
