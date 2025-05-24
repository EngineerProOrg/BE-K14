package services

import (
	"context"
	"social-media/repositories"
	"social-media/services/redisservice"
)

var ctx = context.Background()

type FollowService struct {
	followRepository   repositories.FollowRepository
	redisFollowService *redisservice.RedisFollowService
}

func NewFollowService() *FollowService {
	return &FollowService{
		followRepository:   repositories.NewFollowRepository(),
		redisFollowService: redisservice.NewRedisFollowService(), // assume already initialized
	}
}

func (followService *FollowService) Follow(followerID, followingID int64) error {
	err := followService.followRepository.Follow(followerID, followingID)
	if err != nil {
		return err
	}
	// Redis update
	return followService.redisFollowService.CacheFollow(followerID, followingID)
}

func (s *FollowService) Unfollow(followerID, followingID int64) error {
	err := s.followRepository.Unfollow(followerID, followingID)
	if err != nil {
		return err
	}
	// Redis update
	return s.redisFollowService.RemoveFollowCache(followerID, followingID)
}

func (s *FollowService) GetFollowings(userID int64) ([]int64, error) {
	cached, err := s.redisFollowService.GetCachedFollowings(userID)
	if err == nil && len(cached) > 0 {
		return cached, nil
	}
	// fallback to DB
	followings, err := s.followRepository.GetFollowings(userID)
	if err != nil {
		return nil, err
	}
	_ = s.redisFollowService.CacheFollowingsBulk(userID, followings)
	return followings, nil
}
