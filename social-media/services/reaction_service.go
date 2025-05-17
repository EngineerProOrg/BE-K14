package services

import (
	"social-media/models"
	"social-media/repositories"
)

func GetReactionsByTarget(targetId int64, targetType string) ([]*models.UserReactionResponseViewModel, error) {
	reactionDbModels, err := repositories.GetReactionsByTarget(targetId, targetType)
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

func CountGroupedReactionsByTarget(targetId int64, targetType string) ([]models.ReactionResponseViewModelCount, error) {
	return repositories.CountGroupedReactionsByTarget(targetId, targetType)
}
