package repositories

import (
	"errors"
	"social-media/models"
	"social-media/repositories/databases"
	"time"

	"gorm.io/gorm"
)

func CreateOrUpdateReaction(like *models.Reaction) error {
	var existing models.Reaction

	query := databases.GormDb.Model(&models.Reaction{}).Where("user_id = ?", like.UserId)
	if like.CommentId != nil {
		query = query.Where("comment_id = ?", *like.CommentId)
	} else {
		query = query.Where("post_id = ?", like.PostId).Where("comment_id IS NULL")
	}

	err := query.First(&existing).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// No existing reaction → insert
		return databases.GormDb.Create(like).Error
	}

	if existing.ReactionType == like.ReactionType {
		// Click same icon → remove
		return databases.GormDb.Delete(&existing).Error
	}

	// Different icon → update
	return databases.GormDb.Model(&existing).Updates(map[string]interface{}{
		"reaction_type": like.ReactionType,
		"updated_at":    time.Now(),
	}).Error
}

func DeleteReaction(userId int64, postId int, commentId *int) error {
	query := databases.GormDb.Where("user_id = ?", userId)
	if commentId != nil {
		query = query.Where("comment_id = ?", *commentId)
	} else {
		query = query.Where("post_id = ?", postId).Where("comment_id IS NULL")
	}
	return query.Delete(&models.Reaction{}).Error
}

func CountAllPostLikes() (int64, error) {
	var count int64
	err := databases.GormDb.Model(&models.Reaction{}).Where("post_id IS NOT NULL AND comment_id IS NULL").Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func CountLikesByPostId(postId int64) (int64, error) {
	var count int64
	err := databases.GormDb.Model(&models.Reaction{}).Where("post_id = ? AND comment_id IS NULL", postId).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func CountLikesByCommentId(commentId int64) (int64, error) {
	var count int64
	err := databases.GormDb.Model(&models.Reaction{}).Where("comment_id = ? AND comment_id IS NULL", commentId).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func CountLikes(filter map[string]interface{}) int64 {
	var count int64
	err := databases.GormDb.Model(&models.Reaction{}).Where(filter).Count(&count).Error
	if err != nil {
		return 0
	}
	return count
}

func GetReactionsByPostId(postId int64) ([]models.Reaction, error) {
	var reactions []models.Reaction
	err := databases.GormDb.
		Preload("User").
		Where("post_id = ? AND comment_id IS NULL", postId).
		Find(&reactions).Error
	return reactions, err
}

func CountReactionsByPostId(postId int64) ([]models.ReactionCount, error) {
	var results []models.ReactionCount

	err := databases.GormDb.Model(&models.Reaction{}).
		Select("reaction_type, COUNT(*) as count").
		Where("post_id = ? AND comment_id IS NULL", postId).
		Group("reaction_type").
		Scan(&results).Error
	return results, err
}
