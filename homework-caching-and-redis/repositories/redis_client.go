package repositories

import (
	"context"
	"log"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func InitRedisClient() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})

	// Test ping Redis
	_, err := RedisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("❌ Redis connection failed: %v", err)
	} else {
		log.Println("✅ Redis connected")
	}
}
