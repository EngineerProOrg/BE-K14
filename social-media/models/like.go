package models

import "time"

type Like struct {
	Id        int        `gorm:"primaryKey"`
	PostId    *int       `gorm:"column:post_id;"`    // use pointer to allow nullable
	CommentId *int       `gorm:"column:comment_id;"` // use pointer to allow nullable
	UserId    int        `gorm:"column:user_id;not null"`
	CreatedAt time.Time  `gorm:"column:created_at;not null"`
	UpdatedAt *time.Time `gorm:"column:updated_at;autoUpdateTime:false"`
}
