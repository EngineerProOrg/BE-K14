package repositories

import (
	"errors"
	"social-media/models"
	"social-media/repositories/databases"
	"social-media/utils"
	"time"

	"gorm.io/gorm"
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

func GetCommentById(commentId int64) (*models.Comment, error) {
	commentEntity := &models.Comment{}

	err := databases.GormDb.First(commentEntity, commentId).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, utils.ErrCommentDoesNotExist
	}

	return commentEntity, nil
}

func UpdateComment(comment *models.Comment) (*models.Comment, error) {
	err := databases.GormDb.Model(&models.Comment{}).
		Where("id = ? AND user_id = ? AND post_id = ?", comment.Id, comment.UserId, comment.PostId).
		Updates(map[string]interface{}{
			"updated_at": time.Now(),
			"content":    comment.Content,
		}).Error

	if err != nil {
		return nil, err
	}
	updatedComment, err := GetCommentById(comment.Id)
	return updatedComment, err
}
