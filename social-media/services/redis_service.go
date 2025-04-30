package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"social-media/models/sharedmodels"
	"social-media/repositories/databases"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

func GenerateSessionId(ginContext *gin.Context, userId int64) (string, error) {
	sessionID := uuid.NewString()

	err := databases.RedisClient.Set(ginContext, "session_id:"+sessionID, userId, time.Hour).Err()
	log.Printf("✅ New session created: %s for %d", sessionID, userId)

	if err != nil {
		return "", errors.New("could not store session in Redis")
	}
	return sessionID, nil
}

func SetCachedUserInfo(ginContext *gin.Context, userId int64, author sharedmodels.UserResponseViewModel) {
	userInfoKey := fmt.Sprintf("user_info:%d", userId)

	jsonBytes, err := json.Marshal(author)
	if err != nil {
		log.Printf("❌ Failed to marshal author info: %v", err)
		return
	}

	err = databases.RedisClient.Set(ginContext, userInfoKey, jsonBytes, time.Hour).Err()
	if err != nil {
		log.Printf("❌ Failed to set author info in Redis: %v", err)
	}
}

func GetCachedAuthor(ginContext *gin.Context, userId int64) (sharedmodels.UserResponseViewModel, error) {
	key := fmt.Sprintf("user_info:%d", userId)
	val, err := databases.RedisClient.Get(ginContext, key).Result()
	if err == redis.Nil {
		// fallback to DB when we cannot find userinfo in redis
	}
	var author sharedmodels.UserResponseViewModel
	_ = json.Unmarshal([]byte(val), &author)
	return author, nil
}
