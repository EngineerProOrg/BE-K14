package handlers

import (
	"B5/internal/store"
	"github.com/gin-gonic/gin"
)

type CountHandler struct {
	store store.RedisStore
}

func NewCountHandler(store store.RedisStore) *CountHandler {
	return &CountHandler{store: store}
}

func (h *CountHandler) Handle(c *gin.Context) {
	count, err := h.store.GetUniqueUsersCount(c)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get unique users count"})
		return
	}

	c.JSON(200, gin.H{
		"unique_users_count": count,
	})
}
