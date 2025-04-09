package handlers

import (
	"B5/internal/store"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type LoginHandler struct {
	store store.RedisStore
}

func NewLoginHandler(redisStore store.RedisStore) *LoginHandler {
	return &LoginHandler{store: redisStore}
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
}

func (l *LoginHandler) Handle(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "Username is required"})
		return
	}
	sessionID := uuid.New().String()
	if err := l.store.SaveSession(c, sessionID, req.Username); err != nil {
		c.JSON(500, gin.H{"error": "Failed to create session"})
		return
	}
	c.JSON(200, gin.H{
		"session_id": sessionID,
		"username":   req.Username,
	})
}
