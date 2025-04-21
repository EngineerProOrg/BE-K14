package databases

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func InitRedisClient() *redis.Client {
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
	RedisClient = rdb
	return rdb
}
