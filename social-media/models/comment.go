package models

import (
	"social-media/models/sharedmodels"
	"social-media/utils"
	"time"
)

// db model
type Comment struct {
	Id        int64      `gorm:"primaryKey" json:"id"`
	Content   string     `gorm:"type:text;not null" json:"content"`
	UserId    int64      `gorm:"not null" json:"userId"`
	PostId    int64      `gorm:"not null" json:"postId"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"createdAt"`
	UpdatedAt *time.Time `gorm:"column:updated_at;autoUpdateTime:false"`
	User      User       `gorm:"foreignKey:UserId"`
}

type CommentRequestViewModel struct {
	UserId    int64
	PostId    int64
	CommentId int64
	Content   string `json:"content" binding:"required,notblank"`
	Username  string
}

type CommentResponseViewModel struct {
	Id        int64                          `json:"id"`
	Content   string                         `json:"content"`
	CreatedAt time.Time                      `json:"createdAt"`
	UpdateAt  *time.Time                     `json:"updatedAt"`
	Author    sharedmodels.UserBaseViewModel `json:"author"`
}

func MapCommentRequestViewModelToCommentDbModel(commentRequestViewModel *CommentRequestViewModel) *Comment {
	return &Comment{
		Id:      commentRequestViewModel.CommentId,
		Content: commentRequestViewModel.Content,
		UserId:  commentRequestViewModel.UserId,
		PostId:  commentRequestViewModel.PostId,
	}
}

func (c *Comment) MapCommentEntityAndCommentResponseViewModel(author *UserProfileResponseViewModel) *CommentResponseViewModel {
	return &CommentResponseViewModel{
		Id:        c.Id,
		Content:   c.Content,
		CreatedAt: c.CreatedAt,
		UpdateAt:  c.UpdatedAt,
		Author: sharedmodels.UserBaseViewModel{
			UserId:    author.UserId,
			FirstName: author.FirstName,
			LastName:  author.LastName,
			Name:      author.Name,
			Birthday:  author.Birthday,
			Email:     author.Email,
			Username:  utils.GetUsernameFromEmail(author.Email),
			Avatar:    author.Avatar,
		},
	}
}
