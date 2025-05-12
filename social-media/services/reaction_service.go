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
	for _, reaction := range reactionDbModels {
		userProfileResponseVm := reaction.User.MapUserDbModelToUserProfileResponseViewModel()

		reactionVms = append(reactionVms, reaction.MapReactionDbModelToUserReactionResponseViewModel(userProfileResponseVm))
	}

	return reactionVms, nil
}

func CreateOrUpdateReaction(reactionRequestVm *models.ReactionRequestViewModel) error {
	reactionDbModel := models.MapReactionRequestViewModelToReactionDbModel(reactionRequestVm)
	return repositories.CreateOrUpdateReaction(reactionDbModel)
}
