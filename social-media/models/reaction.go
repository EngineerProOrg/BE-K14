package models

import (
	"social-media/utils"
	"time"
)

// View Models
type ReactionRequestViewModel struct {
	TargetId     int64  `json:"targetId" binding:"required"`
	TargetType   string `json:"targetType" binding:"required,notblank"`
	UserId       int64  `json:"userId"`
	ReactionType string `json:"reactionType" binding:"required,notblank"`
}

type ReactionResponseViewModelCount struct {
	ReactionType string `json:"reactionType"`
	Count        int64  `json:"count"`
}

type UserReactionResponseViewModel struct {
	UserId       int64   `json:"userId"`
	UserName     string  `json:"username"`
	AvatarUrl    *string `json:"avatarUrl"`
	ReactionType string  `json:"reactionType"`
}

// Db models
type Reaction struct {
	Id           int64      `gorm:"primaryKey"`
	TargetId     int64      `gorm:"not null;index:idx_user_target,priority:1"`
	TargetType   string     `gorm:"not null;index:idx_user_target,priority:2"`
	UserId       int64      `gorm:"not null;index:idx_user_target,priority:3"`
	ReactionType string     `gorm:"not null"`
	CreatedAt    time.Time  `gorm:"not null"`
	UpdatedAt    *time.Time `gorm:"autoUpdateTime:false"`

	User User `gorm:"foreignKey:UserId;references:ID"`
}

func (r *Reaction) MapReactionDbModelToUserReactionResponseViewModel(userVm *UserProfileResponseViewModel) *UserReactionResponseViewModel {
	return &UserReactionResponseViewModel{
		UserId:       r.UserId,
		UserName:     utils.GetUsernameFromEmail(userVm.Email),
		AvatarUrl:    userVm.Avatar,
		ReactionType: r.ReactionType,
	}
}

func MapReactionRequestViewModelToReactionDbModel(reactionRequestVm *ReactionRequestViewModel) *Reaction {
	return &Reaction{
		TargetId:     reactionRequestVm.TargetId,
		TargetType:   reactionRequestVm.TargetType,
		UserId:       reactionRequestVm.UserId,
		ReactionType: reactionRequestVm.ReactionType,
		CreatedAt:    time.Now(),
	}
}
