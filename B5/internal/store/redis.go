package store

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

const (
	sessionPrefix  = "session:"
	pingCountKey   = "ping_counts"
	rateLimitKey   = "rate_limit:%s"
	uniqueUsersKey = "unique_users_hll"
)

type RedisStore interface {
	SaveSession(ctx context.Context, sessionID string, username string) error
	GetSession(ctx context.Context, sessionID string) (string, error)
	IncrementPingCount(ctx context.Context, username string) (int64, error)
	CheckRateLimit(ctx context.Context, username string, limit int, window time.Duration) (bool, error)
	GetTopUsers(ctx context.Context, limit int) ([]UserCount, error)
	AddUserToHLL(ctx context.Context, username string) error
	GetUniqueUsersCount(ctx context.Context) (int64, error)
}
type UserCount struct {
	Username string `json:"username"`
	Count    int64  `json:"count"`
}

type redisStore struct {
	client *redis.Client
}

func (r redisStore) AddUserToHLL(ctx context.Context, username string) error {
	return r.client.PFAdd(ctx, uniqueUsersKey, username).Err()
}

func (r redisStore) GetUniqueUsersCount(ctx context.Context) (int64, error) {
	return r.client.PFCount(ctx, uniqueUsersKey).Result()
}

func (r redisStore) IncrementPingCount(ctx context.Context, username string) (int64, error) {
	err := r.client.ZIncrBy(ctx, pingCountKey, 1, username).Err()
	if err != nil {
		return 0, err
	}
	count, err := r.client.ZScore(ctx, pingCountKey, username).Result()
	if err != nil {
		return 0, err
	}
	return int64(count), nil
}

func (r redisStore) SaveSession(ctx context.Context, sessionID string, username string) error {
	return r.client.Set(ctx, sessionPrefix+sessionID, username, time.Hour).Err()
}

func (r redisStore) GetSession(ctx context.Context, sessionID string) (string, error) {
	return r.client.Get(ctx, sessionPrefix+sessionID).Result()
}

func (r redisStore) CheckRateLimit(ctx context.Context, username string, limit int, window time.Duration) (bool, error) {
	key := fmt.Sprintf(rateLimitKey, username)
	count, err := r.client.Incr(ctx, key).Result()
	if err != nil {
		return false, nil
	}
	if count == 1 {
		r.client.Expire(ctx, key, window)
	}
	return count <= int64(limit), nil
}

func (r redisStore) GetTopUsers(ctx context.Context, limit int) ([]UserCount, error) {
	results, err := r.client.ZRevRangeWithScores(ctx, pingCountKey, 0, int64(limit-1)).Result()
	if err != nil {
		return nil, err
	}
	topUsers := make([]UserCount, 0, len(results))
	for _, result := range results {
		topUsers = append(topUsers, UserCount{
			Username: result.Member.(string),
			Count:    int64(result.Score),
		})
	}
	return topUsers, nil
}
func NewRedisStore(addr string, password string, db int) *redisStore {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	return &redisStore{client: client}
}
