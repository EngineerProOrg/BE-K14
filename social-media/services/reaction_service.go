package services

import (
	"social-media/models"
	"social-media/repositories"

	"github.com/gin-gonic/gin"
)

func GetUserReactionsByPostId(ginCtx *gin.Context, postId int64) ([]*models.UserReactionResponseViewModel, error) {
	reactionDbModels, err := repositories.GetReactionsByPostId(postId)
	if err != nil {
		return nil, err
	}

	reactionVms := make([]*models.UserReactionResponseViewModel, 0, len(reactionDbModels))
	for _, like := range reactionDbModels {
		userVm, err := GetCachedUserInfo(ginCtx, like.UserId)

		if err != nil {
			continue
		}

		reactionVms = append(reactionVms, like.MapReactionDbModelToUserReactionResponseViewModel(userVm))
	}

	return reactionVms, nil
}

func CreateOrUpdateReaction(reactionRequestVm *models.ReactionRequestViewModel) error {
	reactionDbModel := models.MapReactionRequestViewModelToReactionDbModel(reactionRequestVm)
	return repositories.CreateOrUpdateReaction(reactionDbModel)
}
