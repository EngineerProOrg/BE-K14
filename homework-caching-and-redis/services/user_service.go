package services

import (
	"errors"
	"homework-caching-and-redis/models"
	"homework-caching-and-redis/repositories"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateANewUser(user *models.User) error {
	return repositories.CreateANewUser(user)
}

func Login(user *models.User) error {
	return repositories.Login(user)
}

func GenerateSessionId(ginContext *gin.Context, email string) (string, error) {
	sessionID := uuid.NewString()

	err := repositories.RedisClient.Set(ginContext, "session:"+sessionID, email, time.Hour).Err()
	if err != nil {
		return "", errors.New("Could not store session in Redis")
	}
	return sessionID, nil
}
