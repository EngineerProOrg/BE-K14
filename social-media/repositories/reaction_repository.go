package repositories

import (
	"errors"
	"social-media/models"
	"social-media/repositories/databases"
	"time"

	"gorm.io/gorm"
)

func CreateOrUpdateReaction(reaction *models.Reaction) error {
	var existing models.Reaction

	err := databases.GormDb.Model(&models.Reaction{}).
		Where("user_id = ? AND target_id = ? AND target_type = ?", reaction.UserId, reaction.TargetId, reaction.TargetType).
		First(&existing).Error

	// if record not found, add new record into db
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return databases.GormDb.Create(reaction).Error
	}

	if existing.ReactionType == reaction.ReactionType {
		return databases.GormDb.Delete(&existing).Error
	}

	// if record found and user changes reaction, updated new reaction
	return databases.GormDb.Model(&existing).Updates(map[string]interface{}{
		"reaction_type": reaction.ReactionType,
		"updated_at":    time.Now(),
	}).Error
}

func DeleteReaction(userId int64, targetId int64, targetType string) error {
	return databases.GormDb.
		Where("user_id = ? AND target_id = ? AND target_type = ?", userId, targetId, targetType).
		Delete(&models.Reaction{}).Error
}

func CountReactions(targetId int64, targetType string) (int64, error) {
	var count int64
	err := databases.GormDb.
		Model(&models.Reaction{}).
		Where("target_id = ? AND target_type = ?", targetId, targetType).
		Count(&count).Error
	return count, err
}

func GetReactionsByTarget(targetId int64, targetType string) ([]models.Reaction, error) {
	var reactions []models.Reaction
	err := databases.GormDb.
		Preload("User").
		Where("target_id = ? AND target_type = ?", targetId, targetType).
		Find(&reactions).Error
	return reactions, err
}

func CountGroupedReactionsByTarget(targetId int64, targetType string) ([]models.ReactionResponseViewModelCount, error) {
	var results []models.ReactionResponseViewModelCount

	err := databases.GormDb.Model(&models.Reaction{}).
		Select("reaction_type, COUNT(*) as count").
		Where("target_id = ? AND target_type = ?", targetId, targetType).
		Group("reaction_type").
		Scan(&results).Error

	return results, err
}
