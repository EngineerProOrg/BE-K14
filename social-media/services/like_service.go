package services

import (
	"social-media/models"
	"social-media/repositories"

	"github.com/gin-gonic/gin"
)

func GetUserReactionsByPostId(ginCtx *gin.Context, postId int64) ([]*models.UserReactionResponseViewModel, error) {
	likes, err := repositories.GetReactionsByPostId(postId)
	if err != nil {
		return nil, err
	}

	reactions := make([]*models.UserReactionResponseViewModel, 0, len(likes))
	for _, like := range likes {
		userVm, err := GetCachedUserInfo(ginCtx, like.UserId)

		if err != nil {
			continue
		}

		reactions = append(reactions, like.MapLikeToUserReactionResponseViewModel(userVm))
	}

	return reactions, nil
}
