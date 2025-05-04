package services

import (
	"social-media/models"
	"social-media/repositories"
	"social-media/utils"
)

func Signup(user *models.User) error {
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	return repositories.Signup(user)
}

func Signin(userSignRequestVm *models.UserSigninRequestViewModel) (*models.UserProfileResponseViewModel, error) {
	userDbModel := models.MapUserSigninRequestViewModelToUserDbModel(userSignRequestVm)
	userModel, err := repositories.Signin(userDbModel)
	if err != nil {
		return nil, err
	}

	userResponse := userModel.MapUserDbModelToUserProfileResponseViewModel()

	return userResponse, err
}

func GetUserProfile(userId int64) (*models.UserProfileResponseViewModel, error) {
	userModel, err := repositories.GetUserProfile(userId)
	if err != nil {
		return nil, err
	}
	userResponse := userModel.MapUserDbModelToUserProfileResponseViewModel()
	return userResponse, err
}

func EditUserProfile(userId int64, vm *models.EditUserProfileRequestViewModel) (*models.UserProfileResponseViewModel, error) {
	updatedUser := models.MapEditUserProfileViewModelToUserDbModel(vm)

	err := repositories.UpdateUserProfile(userId, updatedUser)
	if err != nil {
		return nil, err
	}

	// Get latest user profile from db to cache
	profile, err := GetUserProfile(userId)
	if err != nil {
		return nil, err
	}

	return profile, nil
}
