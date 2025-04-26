package repositories

import (
	"social-media/models"
	"social-media/repositories/databases"
)

func CreatePost(post *models.Post) (*models.Post, error) {
	err := databases.GormDb.Create(post).Error
	if err != nil {
		return nil, err
	}

	err = databases.GormDb.Preload("User").First(post, post.Id).Error
	if err != nil {
		return nil, err
	}
	return post, nil
}
