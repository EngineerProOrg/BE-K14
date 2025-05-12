package models

import (
	"social-media/utils"
	"time"
)

// View Models
type ReactionRequestViewModel struct {
	PostId       int64  `json:"postId" binding:"required,notblank"`
	CommentId    *int64 `json:"commentId"` // allow nullable
	UserId       int64  `json:"user_id"`
	ReactionType string `json:"reactionType" binding:"required,notblank"`
}

type ReactionCount struct {
	ReactionType string
	Count        int64
}

type UserReactionResponseViewModel struct {
	UserId       int64   `json:"user_id"`
	UserName     string  `json:"user_name"`
	AvatarUrl    *string `json:"avatar_url"`
	ReactionType string  `json:"reaction_type"`
}

// Db models
type Reaction struct {
	Id           int64      `gorm:"primaryKey"`
	PostId       int64      `gorm:"index:idx_user_post,unique"`
	CommentId    *int64     `gorm:"index:idx_user_comment,unique"` // for unique constraint with user_id
	UserId       int64      `gorm:"not null;index:idx_user_post,unique;index:idx_user_comment,unique"`
	ReactionType string     `gorm:"column:reaction_type;"`
	CreatedAt    time.Time  `gorm:"column:created_at;not null"`
	UpdatedAt    *time.Time `gorm:"column:updated_at;autoUpdateTime:false"`

	User User `gorm:"foreignKey:UserId"`
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
		PostId:       reactionRequestVm.PostId,
		CommentId:    reactionRequestVm.CommentId,
		UserId:       reactionRequestVm.UserId,
		ReactionType: reactionRequestVm.ReactionType,
		CreatedAt:    time.Now(),
	}
}
