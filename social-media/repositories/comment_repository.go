package repositories

import (
	"social-media/models"
	"social-media/repositories/databases"
)

func CreateComment(comment *models.Comment) (*models.Comment, error) {
	err := databases.GormDb.Create(comment).Error
	if err != nil {
		return nil, err
	}
	return comment, nil
}

func GetCommentsByPostId(postId int64) ([]models.Comment, error) {
	var comments []models.Comment
	err := databases.GormDb.Preload("User").Where("post_id = ?", postId).Order("created_at ASC").Find(&comments).Error
	if err != nil {
		return nil, err
	}
	return comments, nil
}
