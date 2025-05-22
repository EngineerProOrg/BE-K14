package redisservice

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"social-media/repositories/databases"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type RedisFollowService struct {
	rdb *redis.Client
}

func NewFollowRedisService() *RedisFollowService {
	return &RedisFollowService{
		rdb: databases.RedisClient,
	}
}

func (r *RedisFollowService) CacheFollow(followerID, followingID int64) error {
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

func (r *RedisFollowService) RemoveFollowCache(followerID, followingID int64) error {
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

func (r *RedisFollowService) GetCachedFollowings(userID int64) ([]int64, error) {
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

func (r *RedisFollowService) CacheFollowingsBulk(userID int64, ids []int64) error {
	key := fmt.Sprintf("user:%d:following", userID)
	for _, id := range ids {
		r.rdb.SAdd(ctx, key, id)
	}
	r.rdb.Expire(ctx, key, time.Hour*6)
	return nil
}
