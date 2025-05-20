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
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type RedisService struct {
	rdb *redis.Client
}

func NewRedisService() *RedisService {
	return &RedisService{
		rdb: databases.RedisClient,
	}
}

// Set cached user info by username
func SetCachedUserInfoByUsername(username string, userProfileResponseVm *models.UserProfileResponseViewModel) {
	ctx := context.Background()
	key := fmt.Sprintf("user_info:%s", username)

	jsonBytes, err := json.Marshal(userProfileResponseVm)
	if err != nil {
		log.Printf("❌ Failed to marshal user info: %v", err)
		return
	}

	err = databases.RedisClient.Set(ctx, key, jsonBytes, time.Hour).Err()
	if err != nil {
		log.Printf("❌ Failed to set user info in Redis: %v", err)
	}
}

// Get cached user info, fallback to DB if not found
func GetCachedUserInfoByUsername(username string) (*models.UserProfileResponseViewModel, error) {
	ctx := context.Background()
	key := fmt.Sprintf("user_info:%s", username)

	val, err := databases.RedisClient.Get(ctx, key).Result()
	if err != nil {
		// Redis miss, fallback to DB
		log.Printf("❌ Failed to get from Redis: %v", err)
		userProfile, err := GetUserProfileByUsername(username)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, utils.ErrUserDoesNotExist
			}
			return nil, fmt.Errorf("failed to query user from DB: %w", err)
		}
		return userProfile, nil
	}

	var profile models.UserProfileResponseViewModel
	if err := json.Unmarshal([]byte(val), &profile); err != nil {
		return nil, fmt.Errorf("failed to unmarshal cached user info: %w", err)
	}

	return &profile, nil
}

// Check if username exists in Redis
func CheckUsernameExistInRedis(username string) bool {
	ctx := context.Background()
	key := fmt.Sprintf("user_info:%s", username)
	_, err := databases.RedisClient.Get(ctx, key).Result()
	return err == nil
}

// Get user ID from Redis by username
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

// Delete cached user info
func DeleteCachedUserInfo(username string) error {
	ctx := context.Background()
	key := fmt.Sprintf("user_info:%s", username)

	err := databases.RedisClient.Del(ctx, key).Err()
	if err != nil {
		log.Printf("❌ Failed to delete cached user info for username=%s: %v", username, err)
		return err
	}

	return nil
}

func (r *RedisService) CacheFollow(followerID, followingID int64) error {
	fkey := fmt.Sprintf("user:%d:following", followerID)
	bkey := fmt.Sprintf("user:%d:followers", followingID)
	if err := r.rdb.SAdd(ctx, fkey, followingID).Err(); err != nil {
		return err
	}
	if err := r.rdb.SAdd(ctx, bkey, followerID).Err(); err != nil {
		return err
	}
	r.rdb.Expire(ctx, fkey, time.Hour*6)
	r.rdb.Expire(ctx, bkey, time.Hour*6)
	return nil
}

func (r *RedisService) RemoveFollowCache(followerID, followingID int64) error {
	fkey := fmt.Sprintf("user:%d:following", followerID)
	bkey := fmt.Sprintf("user:%d:followers", followingID)
	if err := r.rdb.SRem(ctx, fkey, followingID).Err(); err != nil {
		return err
	}
	if err := r.rdb.SRem(ctx, bkey, followerID).Err(); err != nil {
		return err
	}
	return nil
}

func (r *RedisService) GetCachedFollowings(userID int64) ([]int64, error) {
	key := fmt.Sprintf("user:%d:following", userID)
	ids, err := r.rdb.SMembers(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	var result []int64
	for _, idStr := range ids {
		id, _ := strconv.ParseInt(idStr, 10, 64)
		result = append(result, id)
	}
	return result, nil
}

func (r *RedisService) CacheFollowingsBulk(userID int64, ids []int64) error {
	key := fmt.Sprintf("user:%d:following", userID)
	for _, id := range ids {
		r.rdb.SAdd(ctx, key, id)
	}
	r.rdb.Expire(ctx, key, time.Hour*6)
	return nil
}
