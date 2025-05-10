package repositories

import (
	"social-media/models"
	"social-media/repositories/databases"

	"gorm.io/gorm"
)

// basePostQuery returns a GORM query builder for posts.
// If includeAuthorInfo is true, it will preload the associated User (author) data.
// Use this to avoid repeating Preload("User") across repo methods.
func basePostQuery(includeAuthorInfo bool) *gorm.DB {
	query := databases.GormDb.Model(&models.Post{})
	if includeAuthorInfo {
		query = query.Preload("User")
	}
	return query
}
