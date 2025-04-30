package models

import (
	"social-media/models/sharedmodels"
	"time"
)

type PostRequestViewModel struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
	UserId  int64  `json:"userId"`
}

type PostResponseViewModel struct {
	PostId    int                            `json:"postId"`
	Title     string                         `json:"title"`
	Content   string                         `json:"content"`
	CreatedAt time.Time                      `json:"createdAt"`
	Author    sharedmodels.UserBaseViewModel `json:"author"`
}

// Db model
type Post struct {
	Id        int        `gorm:"primaryKey;column:id"`
	Title     string     `gorm:"column:title;size:500;not null"`
	Content   string     `gorm:"column:content;type:text"`
	UserId    int64      `gorm:"column:user_id;not null"`
	User      User       `gorm:"foreignKey:UserId"`
	CreatedAt time.Time  `gorm:"column:created_at;not null"`
	UpdatedAt *time.Time `gorm:"column:updated_at;autoUpdateTime:false"`
}

func CreateMappingPostRequestViewModelToPostEntity(vm *PostRequestViewModel) *Post {
	return &Post{
		Title:     vm.Title,
		Content:   vm.Content,
		UserId:    vm.UserId,
		CreatedAt: time.Now(),
	}
}

func (p *Post) CreateMappingPostEntityToPostResponseViewModel() *PostResponseViewModel {
	return &PostResponseViewModel{
		PostId:    p.Id,
		Title:     p.Title,
		Content:   p.Content,
		CreatedAt: p.CreatedAt,
		Author: sharedmodels.UserBaseViewModel{
			UserId:    p.UserId,
			FirstName: p.User.FirstName,
			LastName:  p.User.LastName,
			Name:      p.User.Name,
			Birthday:  p.User.Birthday,
			Email:     p.User.Email,
			Avatar:    p.User.Avatar,
		},
	}
}
