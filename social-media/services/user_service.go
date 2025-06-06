package services

import (
	"social-media/models"
	"social-media/repositories"
	"social-media/utils"

	"github.com/gin-gonic/gin"
)

func Signup(gctx *gin.Context, userSignupRequestVm *models.UserSignupRequestViewModel) (*models.UserProfileResponseViewModel, error) {
	hashedPassword, err := utils.HashPassword(userSignupRequestVm.Password)
	if err != nil {
		return nil, err
	}

	userdb := models.MapUserSignupRequestViewModelToUserDbModel(userSignupRequestVm)
	userdb.Password = hashedPassword

	userdb, err = repositories.Signup(userdb)

	if err != nil {
		return nil, err
	}

	userProfileResponseViewModel := userdb.MapUserDbModelToUserProfileResponseViewModel()

	// After signup successfully, we should cache user info into db
	SetCachedUserInfoByUsername(userdb.Username, userProfileResponseViewModel)
	return userProfileResponseViewModel, nil
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

func GetUserProfile(userid int64) (*models.UserProfileResponseViewModel, error) {
	userModel, err := repositories.GetUserProfile(userid)
	if err != nil {
		return nil, err
	}
	userResponse := userModel.MapUserDbModelToUserProfileResponseViewModel()
	return userResponse, err
}

func GetUserProfileByUsername(username string) (*models.UserProfileResponseViewModel, error) {
	userModel, err := repositories.GetUserProfileByUsername(username)
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
