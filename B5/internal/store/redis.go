package store

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type RedisStore interface {
	SaveSession(ctx context.Context, sessionID string, username string) error
	GetSession(ctx context.Context, sessionID string) (string, error)
}

type redisStore struct {
	client *redis.Client
}

func (r redisStore) SaveSession(ctx context.Context, sessionID string, username string) error {
	return r.client.Set(ctx, "session:"+sessionID, username, time.Hour).Err()
}

func (r redisStore) GetSession(ctx context.Context, sessionID string) (string, error) {
	return r.client.Get(ctx, "session:"+sessionID).Result()
}

func NewRedisStore() *redisStore {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	return &redisStore{client: client}
}
