package store

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisStore interface {
	SaveSession(ctx context.Context, sessionID string, username string) error
	GetSession(ctx context.Context, sessionID string) (string, error)
	IncrementPingCount(ctx context.Context, username string) (int64, error)
	CheckRateLimit(ctx context.Context, username string) (bool, error)
}

type redisStore struct {
	client *redis.Client
}

func (r redisStore) IncrementPingCount(ctx context.Context, username string) (int64, error) {
	key := "ping_count" + username
	return r.client.Incr(ctx, key).Result()
}

func (r redisStore) SaveSession(ctx context.Context, sessionID string, username string) error {
	return r.client.Set(ctx, "session:"+sessionID, username, time.Hour).Err()
}

func (r redisStore) GetSession(ctx context.Context, sessionID string) (string, error) {
	return r.client.Get(ctx, "session:"+sessionID).Result()
}

func (r redisStore) CheckRateLimit(ctx context.Context, username string) (bool, error) {
	key := fmt.Sprintf("rate_limit:%s", username)
	count, err := r.client.Incr(ctx, key).Result()
	if err != nil {
		return false, nil
	}
	if count == 1 {
		r.client.Expire(ctx, key, 60*time.Second)
	}
	return count <= 2, nil
}

func NewRedisStore() *redisStore {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return &redisStore{client: client}
}
