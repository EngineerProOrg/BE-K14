package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"social-media/models"
	"social-media/repositories/databases"
	"social-media/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
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

func SetCachedUserInfoByUsername(ginContext *gin.Context, username string, userProfileResponseVm *models.UserProfileResponseViewModel) {
	userInfoKey := fmt.Sprintf("user_info:%s", username)

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

// Cached User Info inside Redis. Therefore, we don't need to call query.
func GetCachedUserInfoByUsername(ginContext *gin.Context, username string) (*models.UserProfileResponseViewModel, error) {
	key := fmt.Sprintf("user_info:%s", username)
	val, err := databases.RedisClient.Get(ginContext, key).Result()

	if err != nil {
		// Redis could't not find any user then fallback to query db
		log.Printf("❌ Failed to get from Redis: %v", err)
		userSigninResponseVm, err := GetUserProfileByUsername(username)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				// cannot find any user in db, return not found.
				return nil, utils.ErrUserDoesNotExist
			}
			// something else network connection
			return nil, fmt.Errorf("failed to query user from DB: %w", err)
		}
		return userSigninResponseVm, nil
	}

	var userSigninResponseVm models.UserProfileResponseViewModel
	if err := json.Unmarshal([]byte(val), &userSigninResponseVm); err != nil {
		return nil, fmt.Errorf("failed to unmarshal cached user info: %w", err)
	}

	return &userSigninResponseVm, nil
}

func CheckUsernameExistInRedis(ginContext *gin.Context, username string) bool {
	key := fmt.Sprintf("user_info:%s", username)
	_, err := databases.RedisClient.Get(ginContext, key).Result()
	return err == nil
}

func GetUserIdByUsernameFromRedis(username string) (int64, error) {
	ctx := context.Background()
	key := fmt.Sprintf("user_info:%s", username)
	val, err := databases.RedisClient.Get(ctx, key).Result()
	if err != nil {
		return 0, err
	}

	var profile models.UserProfileResponseViewModel
	if err := json.Unmarshal([]byte(val), &profile); err != nil {
		return 0, err
	}

	return profile.UserId, nil
}

func DeleteCachedUserInfo(ctx *gin.Context, username string) error {
	userInfoKey := fmt.Sprintf("user_info:%s", username)
	err := databases.RedisClient.Del(ctx, userInfoKey).Err()
	if err != nil {
		log.Printf("❌ Failed to delete cached user info for username=%s: %v", username, err)
		return err
	}
	return nil
}
