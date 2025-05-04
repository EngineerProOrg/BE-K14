package models

import (
	"social-media/models/sharedmodels"
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
	Content string `json:"content" binding:"required,notblank"`
}

type CommentResponseViewModel struct {
	Id        int64                          `json:"id"`
	Content   string                         `json:"content"`
	CreatedAt time.Time                      `json:"createdAt"`
	UpdateAt  *time.Time                     `json:"updatedAt"`
	Author    sharedmodels.UserBaseViewModel `json:"author"`
}

func (c *Comment) CreateMappingCommentEntityAndCommentResponseViewModel() *CommentResponseViewModel {
	return &CommentResponseViewModel{
		Id:        c.Id,
		Content:   c.Content,
		CreatedAt: c.CreatedAt,
		UpdateAt:  c.UpdatedAt,
		Author: sharedmodels.UserBaseViewModel{
			UserId:    c.UserId,
			FirstName: c.User.FirstName,
			LastName:  c.User.LastName,
			Name:      c.User.Name,
			Birthday:  c.User.Birthday,
			Email:     c.User.Email,
			Avatar:    c.User.Avatar,
		},
	}
}
