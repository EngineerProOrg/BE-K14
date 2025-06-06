package models

import "time"

// Db model
type Follow struct {
	Id          int64      `gorm:"primaryKey;column:id"`
	FollowerId  int        `gorm:"column:follower_id;not null"`
	FollowingId int        `gorm:"column:following_id;not null"`
	CreatedAt   time.Time  `gorm:"column:created_at;not null"`
	UpdatedAt   *time.Time `gorm:"column:updated_at;autoUpdateTime:false"` //  *time.Time allows nullable
}
