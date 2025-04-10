package handlers

import (
	"B5/internal/store"
	"github.com/gin-gonic/gin"
	"sync"
	"time"
)

type PingHandler struct {
	mu    sync.Mutex
	store store.RedisStore
}

func NewPingHandler(store store.RedisStore) *PingHandler {
	return &PingHandler{store: store}
}

func (p *PingHandler) Handle(c *gin.Context) {
	sessionID := c.GetHeader("X-Session-ID")
	if sessionID == "" {
		c.JSON(401, gin.H{"error": "No session ID provided"})
		return
	}
	username, err := p.store.GetSession(c, sessionID)
	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid Session"})
		return
	}
	allowed, err := p.store.CheckRateLimit(c, username)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to check rate limit"})
		return
	}
	if !allowed {
		c.JSON(429, gin.H{"error": "Rate limit exceeded. Try again later."})
		return
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	count, err := p.store.IncrementPingCount(c, username)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to increment ping count"})
		return
	}
	time.Sleep(5 * time.Second)
	c.JSON(200, gin.H{"message": "pong", "count": count})
}
