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
	p.mu.Lock()
	defer p.mu.Unlock()
	time.Sleep(5 * time.Second)
	c.JSON(200, gin.H{"message": "pong"})
}
