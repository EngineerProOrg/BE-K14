package middleware

import (
	"B5/internal/store"
	"github.com/gin-gonic/gin"
	"time"
)

type RateLimiter struct {
	store  store.RedisStore
	limit  int
	window time.Duration
}

func NewRateLimiter(store store.RedisStore, limit int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		store:  store,
		limit:  limit,
		window: window,
	}
}

func (r *RateLimiter) Limit() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.GetString("username")
		allowed, err := r.store.CheckRateLimit(c, username, r.limit, r.window)
		if err != nil {
			c.JSON(500, gin.H{"errors": "Rate limit check failed"})
			c.Abort()
			return
		}
		if !allowed {
			c.JSON(429, gin.H{"errors": "Rate limit exceeded"})
			c.Abort()
			return
		}
		c.Next()
	}
}
