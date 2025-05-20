package services

import (
	"context"
	"social-media/repositories"
)

var ctx = context.Background()

type FollowService struct {
	followRepository repositories.FollowRepository
	redisService     *RedisService
}

func NewFollowService() *FollowService {
	return &FollowService{
		followRepository: repositories.NewFollowRepository(),
		redisService:     NewRedisService(), // assume already initialized
	}
}

func (s *FollowService) Follow(followerID, followingID int64) error {
	err := s.followRepository.Follow(followerID, followingID)
	if err != nil {
		return err
	}
	// Redis update
	return s.redisService.CacheFollow(followerID, followingID)
}

func (s *FollowService) Unfollow(followerID, followingID int64) error {
	err := s.followRepository.Unfollow(followerID, followingID)
	if err != nil {
		return err
	}
	// Redis update
	return s.redisService.RemoveFollowCache(followerID, followingID)
}

func (s *FollowService) GetFollowings(userID int64) ([]int64, error) {
	cached, err := s.redisService.GetCachedFollowings(userID)
	if err == nil && len(cached) > 0 {
		return cached, nil
	}
	// fallback to DB
	followings, err := s.followRepository.GetFollowings(userID)
	if err != nil {
		return nil, err
	}
	_ = s.redisService.CacheFollowingsBulk(userID, followings)
	return followings, nil
}
