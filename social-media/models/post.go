package models

import (
	"social-media/models/sharedmodels"
	"time"
)

type PostRequestViewModel struct {
	Title   string `json:"title" binding:"required,notblank"`
	Content string `json:"content" binding:"required,notblank"`
	UserId  int64  `json:"userId"`
}

type PostWithAuthorViewModel struct {
	Post   sharedmodels.PostBaseViewModel `json:"post"`
	Author sharedmodels.UserBaseViewModel `json:"author"`
}

type PostsWithAuthorResponse struct {
	Posts []PostWithAuthorViewModel `json:"posts"`
}

type PostResponseViewModel struct {
	Post sharedmodels.PostBaseViewModel `json:"post"`
}

type CreatedOrUpdatedPostResponseViewModel struct {
	Author UserProfileResponseViewModel   `json:"author"`
	Post   sharedmodels.PostBaseViewModel `json:"post"`
}

type PostUserResponseViewModel struct {
	Author UserProfileResponseViewModel `json:"author"`
	Posts  []PostResponseViewModel      `json:"posts"`
}

// Db model
type Post struct {
	Id        int64      `gorm:"primaryKey;column:id"`
	Title     string     `gorm:"column:title;size:500;not null"`
	Content   string     `gorm:"column:content;type:text"`
	UserId    int64      `gorm:"column:user_id;not null"`
	User      User       `gorm:"foreignKey:UserId"`
	CreatedAt time.Time  `gorm:"column:created_at;not null"`
	UpdatedAt *time.Time `gorm:"column:updated_at;autoUpdateTime:false"`
}

func MapPostRequestViewModelToPostDbModel(vm *PostRequestViewModel) *Post {
	return &Post{
		Title:     vm.Title,
		Content:   vm.Content,
		UserId:    vm.UserId,
		CreatedAt: time.Now(),
	}
}

func (p *Post) MapPostDbModelToPostResponseViewModel() *PostResponseViewModel {
	return &PostResponseViewModel{
		Post: sharedmodels.PostBaseViewModel{
			PostId:    p.Id,
			Title:     p.Title,
			Content:   p.Content,
			CreatedAt: p.CreatedAt,
			UpdateAt:  p.UpdatedAt,
		},
	}
}

func (p *Post) MapPostDbModelToCreatedPostResponseViewModel(author *UserProfileResponseViewModel) *CreatedOrUpdatedPostResponseViewModel {
	return &CreatedOrUpdatedPostResponseViewModel{
		Author: *author,
		Post: sharedmodels.PostBaseViewModel{
			PostId:    p.Id,
			Title:     p.Title,
			Content:   p.Content,
			CreatedAt: p.CreatedAt,
			UpdateAt:  p.UpdatedAt,
		},
	}
}

func (p *Post) MapPostDbModelToPostWithAuthorViewModel() *PostWithAuthorViewModel {
	return &PostWithAuthorViewModel{
		Post: sharedmodels.PostBaseViewModel{
			PostId:    p.Id,
			Title:     p.Title,
			Content:   p.Content,
			CreatedAt: p.CreatedAt,
			UpdateAt:  p.UpdatedAt,
		},
		Author: sharedmodels.UserBaseViewModel{
			UserId:    p.UserId,
			FirstName: p.User.FirstName,
			LastName:  p.User.LastName,
			Name:      p.User.Name,
			Birthday:  p.User.Birthday,
			Email:     p.User.Email,
			Username:  p.User.Username,
			Avatar:    p.User.Avatar,
		},
	}
}
