package models

import "time"

type PostViewModel struct {
	Title     string    `json:"title" binding:"required"`
	Content   string    `json:"content" binding:"required"`
	UserId    int       `json:"userId" binding:"required"`
	CreatedAt time.Time `json:"createdAt" binding:"required"`
}

// Db model
type Post struct {
	Id        int        `gorm:"primaryKey;column:id"`
	Title     string     `gorm:"column:title;size:500;not null"`
	Content   string     `gorm:"column:content;type:text"`
	UserId    int        `gorm:"column:user_id;not null"`
	User      User       `gorm:"foreignKey:UserId"`
	CreatedAt time.Time  `gorm:"column:created_at;not null"`
	UpdatedAt *time.Time `gorm:"column:updated_at;autoUpdateTime:false"`
}

func CreateMappingPostViewModelToPostEntity(vm *PostViewModel) *Post {
	return &Post{
		Title:     vm.Title,
		Content:   vm.Content,
		UserId:    vm.UserId,
		CreatedAt: time.Now(),
	}
}
