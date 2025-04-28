package services

import (
	"errors"
	"log"
	"social-media/repositories/databases"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GenerateSessionId(ginContext *gin.Context, userId int64) (string, error) {
	sessionID := uuid.NewString()

	err := databases.RedisClient.Set(ginContext, "session_id:"+sessionID, userId, time.Hour).Err()
	log.Printf("✅ New session created: %s for %s", sessionID, userId)

	if err != nil {
		return "", errors.New("could not store session in Redis")
	}
	return sessionID, nil
}
