package middleware

import (
	"B5/internal/store"
	"github.com/gin-gonic/gin"
)

type SessionMiddleware struct {
	store store.RedisStore
}

func NewSessionMiddleware(store store.RedisStore) *SessionMiddleware {
	return &SessionMiddleware{store: store}
}

func (m *SessionMiddleware) ValidateSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionID := c.GetHeader("X-Session-ID")
		if sessionID == "" {
			c.JSON(401, gin.H{"error": "No session ID provided"})
			c.Abort()
			return
		}
		username, err := m.store.GetSession(c, sessionID)
		if err != nil {
			c.JSON(401, gin.H{"error": "Invalid Session"})
			c.Abort()
			return
		}
		c.Set("username", username)
		c.Next()
	}
}
