package databases

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"social-media/models"
	"time"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func InitRedisClient() {
	addr := os.Getenv("REDIS_ADDR") //Example: localhost:6379
	if addr == "" {
		addr = "localhost:6379"
	}

	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
		DB:   0,
	})

	err := rdb.Ping(context.Background()).Err()
	if err != nil {
		log.Fatalf("❌ Cannot connect to Redis: %v", err)
	}

	fmt.Println("✅ Connected to Redis successfully!")
	setupDefaultCache(rdb)

	RedisClient = rdb
}

func setupDefaultCache(redisClient *redis.Client) {
	var users []models.User
	err := GormDb.Find(&users).Error
	if err != nil {
		log.Fatalf("❌ Failed to query users for caching: %v\n", err)
		return
	}

	ctx := context.Background()

	for _, user := range users {
		// Convert user model to response view model
		profile := user.MapUserDbModelToUserProfileResponseViewModel()

		// Marshal profile to JSON
		jsonValue, err := json.Marshal(profile)
		if err != nil {
			log.Printf("❌ Failed to marshal profile of %s: %v\n", user.Username, err)
			continue
		}

		// Cache in Redis
		key := fmt.Sprintf("user_info:%s", user.Username)
		err = redisClient.Set(ctx, key, jsonValue, 24*time.Hour).Err()
		if err != nil {
			log.Printf("❌ Failed to cache profile for %s: %v\n", user.Username, err)
		} else {
			fmt.Printf("✅ Cached profile of %s in Redis\n", user.Username)
		}
	}
}
