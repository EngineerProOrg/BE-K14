package models

import (
	"social-media/utils"
	"time"
)

// Db models
type Like struct {
	Id           int        `gorm:"primaryKey"`
	PostId       *int       `gorm:"column:post_id;"`    // use pointer to allow nullable
	CommentId    *int       `gorm:"column:comment_id;"` // use pointer to allow nullable
	UserId       int64      `gorm:"column:user_id;not null"`
	ReactionType string     `gorm:"column:reaction_type;"`
	CreatedAt    time.Time  `gorm:"column:created_at;not null"`
	UpdatedAt    *time.Time `gorm:"column:updated_at;autoUpdateTime:false"`
}

func (like *Like) MapLikeToUserReactionResponseViewModel(userVm *UserProfileResponseViewModel) *UserReactionResponseViewModel {
	return &UserReactionResponseViewModel{
		UserId:       like.UserId,
		UserName:     utils.GetUsernameFromEmail(userVm.Email),
		AvatarUrl:    userVm.Avatar,
		ReactionType: like.ReactionType,
	}
}
