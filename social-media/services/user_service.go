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

func Signin(userInput *models.User) error {
	return repositories.Signin(userInput)
}

func GetUserProfile(userId int64) (*models.User, error) {
	return repositories.GetUserProfile(userId)
}

func EditUserProfile(userId int64, vm *models.EditUserProfileViewModel) error {
	updatedUser := models.MapEditUserProfileViewModelToUserEntity(vm)

	return repositories.UpdateUserProfile(userId, updatedUser)
}
