package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"social-media/models"
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

func SetCachedUserInfo(ginContext *gin.Context, userId int64, userProfileResponseVm *models.UserProfileResponseViewModel) {
	userInfoKey := fmt.Sprintf("user_info:%d", userId)

	jsonBytes, err := json.Marshal(userProfileResponseVm)
	if err != nil {
		log.Printf("❌ Failed to marshal author info: %v", err)
		return
	}

	err = databases.RedisClient.Set(ginContext, userInfoKey, jsonBytes, time.Hour).Err()
	if err != nil {
		log.Printf("❌ Failed to set author info in Redis: %v", err)
	}
}

func GetCachedUserInfo(ginContext *gin.Context, userId int64) (*models.UserProfileResponseViewModel, error) {
	key := fmt.Sprintf("user_info:%d", userId)
	val, err := databases.RedisClient.Get(ginContext, key).Result()

	if err != nil {
		if err == redis.Nil {
			return nil, fmt.Errorf("user not found in Redis")
		}
		return nil, fmt.Errorf("failed to get from Redis: %w", err)
	}

	userSigninResponseVm := &models.UserProfileResponseViewModel{}
	_ = json.Unmarshal([]byte(val), &userSigninResponseVm)
	return userSigninResponseVm, nil
}

func DeleteCachedUserInfo(ctx *gin.Context, userId int64) error {
	userInfoKey := fmt.Sprintf("user_info:%d", userId)
	err := databases.RedisClient.Del(ctx, userInfoKey).Err()
	if err != nil {
		log.Printf("❌ Failed to delete cached user info for userId=%d: %v", userId, err)
		return err
	}
	return nil
}
