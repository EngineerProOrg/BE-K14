package services

import (
	"social-media/models"
	"social-media/repositories"
)

func CreatePost(post *models.Post) (*models.Post, error) {
	return repositories.CreatePost(post)
}
