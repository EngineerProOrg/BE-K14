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

func Signin(userInput *models.User) (*models.UserSigninResponseViewModel, error) {
	userModel, err := repositories.Signin(userInput)
	userResponse := &models.UserSigninResponseViewModel{
		FirstName: userModel.FirstName,
		LastName:  userModel.LastName,
		Birthday:  userModel.Birthday,
		Email:     userModel.Email,
		UserId:    userModel.ID,
	}
	return userResponse, err
}

func GetUserProfile(userId int64) (*models.User, error) {
	return repositories.GetUserProfile(userId)
}

func EditUserProfile(userId int64, vm *models.EditUserProfileViewModel) error {
	updatedUser := models.MapEditUserProfileViewModelToUserEntity(vm)

	return repositories.UpdateUserProfile(userId, updatedUser)
}
