package repositories

import (
	"social-media/models"
	"social-media/repositories/databases"
)

type FollowRepository struct{}

func NewFollowRepository() FollowRepository {
	return FollowRepository{}
}

func (fr FollowRepository) Follow(followerId, followingId int64) error {
	follow := models.Follow{
		FollowerId:  int(followerId),
		FollowingId: int(followingId),
	}
	return databases.GormDb.Create(&follow).Error
}

func (r FollowRepository) Unfollow(followerID, followingID int64) error {
	return databases.GormDb.Where("follower_id = ? AND following_id = ?", followerID, followingID).Delete(&models.Follow{}).Error
}

func (r FollowRepository) GetFollowings(userID int64) ([]int64, error) {
	var follows []models.Follow
	err := databases.GormDb.Where("follower_id = ?", userID).Find(&follows).Error
	if err != nil {
		return nil, err
	}
	ids := make([]int64, 0, len(follows))
	for _, f := range follows {
		ids = append(ids, int64(f.FollowingId))
	}
	return ids, nil
}
