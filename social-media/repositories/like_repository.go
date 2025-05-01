package repositories

import (
	"social-media/models"
	"social-media/repositories/databases"
)

func LikePostOrComment(like *models.Like) error {
	err := databases.GormDb.Create(like).Error
	return err
}

func UnlikePostOrComment(like *models.Like) error {
	err := databases.GormDb.Delete(like).Error
	return err
}

func CountAllPostLikes() (int64, error) {
	var count int64
	err := databases.GormDb.Model(&models.Like{}).Where("post_id IS NOT NULL AND comment_id IS NULL").Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func CountLikesByPostId(postId int64) (int64, error) {
	var count int64
	err := databases.GormDb.Model(&models.Like{}).Where("post_id = ? AND comment_id IS NULL", postId).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func CountLikesByCommentId(commentId int64) (int64, error) {
	var count int64
	err := databases.GormDb.Model(&models.Like{}).Where("comment_id = ? AND comment_id IS NULL", commentId).Count(&count).Error
	if err != nil {
		return 0, err
	}
	return count, nil
}

func CountLikes(filter map[string]interface{}) int64 {
	var count int64
	err := databases.GormDb.Model(&models.Like{}).Where(filter).Count(&count).Error
	if err != nil {
		return 0
	}
	return count
}

func GetReactionsByPostId(postId int64) ([]models.Like, error) {
	var likes []models.Like
	err := databases.GormDb.
		Where("post_id = ? AND comment_id IS NULL", postId).
		Find(&likes).Error
	return likes, err
}

func CountReactionsByPostId(postId int64) ([]models.ReactionCount, error) {
	var results []models.ReactionCount

	err := databases.GormDb.Model(&models.Like{}).
		Select("reaction_type, COUNT(*) as count").
		Where("post_id = ? AND comment_id IS NULL", postId).
		Group("reaction_type").
		Scan(&results).Error
	return results, err
}
