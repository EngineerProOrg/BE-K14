package controllers

import (
	"errors"
	"fmt"
	"net/http"
	"social-media/models"
	"social-media/services"
	"social-media/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Signup(context *gin.Context) {
	// 1) Validate and bind request body
	userSignupViewModel, ok := utils.BindAndValidate[models.UserSignupRequestViewModel](context)
	if !ok {
		return
	}

	// 2) Priorize to get user info from cache. If exist return 409
	ok = services.CheckUsernameExistInRedis(utils.GetUsernameFromEmail(userSignupViewModel.Email))
	if ok {
		errorMessage := fmt.Sprintf("email %s has been registered", userSignupViewModel.Email)
		context.JSON(http.StatusConflict, gin.H{"error": errorMessage})
		return
	}

	userProfileResponseViewModel, err := services.Signup(context, userSignupViewModel)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"userProfile": userProfileResponseViewModel})
}

func Signin(context *gin.Context) {
	userInput, ok := utils.BindAndValidate[models.UserSigninRequestViewModel](context)
	if !ok {
		return
	}

	// Priorize to get user info from cache
	ok = services.CheckUsernameExistInRedis(utils.GetUsernameFromEmail(userInput.Email))
	if !ok {
		context.JSON(http.StatusNotFound, gin.H{"error": utils.ErrInvalidLogin.Error()})
		return
	}

	userSigninResponseVm, err := services.Signin(userInput)
	if err != nil {
		if errors.Is(err, utils.ErrInvalidLogin) {
			context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	accessToken, err := utils.GenerateAccessToken(userInput.Email, userSigninResponseVm.UserId)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"accessToken": accessToken, "userInfo": userSigninResponseVm})
}

func GetUserProfile(context *gin.Context) {
	userId, err := strconv.ParseInt(context.Param("userId"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "userId": userId})
		return
	}

	extractedUserId, ok := ExtractUserIdFromAccessToken(context)
	if !ok {
		return
	}

	if extractedUserId != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	extractedUsername, ok := ExtractUsernameFromAccessToken(context)
	if !ok {
		return
	}

	userSigninResponseVm, err := services.GetCachedUserInfoByUsername(extractedUsername)
	if err != nil {
		if errors.Is(err, utils.ErrUserDoesNotExist) {
			context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	context.JSON(http.StatusOK, gin.H{"userProfile": userSigninResponseVm})
}

func EditUserProfile(context *gin.Context) {
	userId, err := strconv.ParseInt(context.Param("userId"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "userId": userId})
		return
	}

	extractedUserId, ok := ExtractUserIdFromAccessToken(context)
	if !ok {
		return
	}

	if extractedUserId != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	_, err = services.GetUserProfile(userId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error(), "userId": userId})
		return
	}

	editUserProfile, ok := utils.BindAndValidate[models.EditUserProfileRequestViewModel](context)
	if !ok {
		return
	}

	updatedProfile, err := services.EditUserProfile(userId, editUserProfile)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "userId": userId})
		return
	}
	// Set cached
	services.SetCachedUserInfoByUsername(utils.GetUsernameFromEmail(updatedProfile.Email), updatedProfile)
	context.JSON(http.StatusOK, gin.H{"userInfo": updatedProfile})
}

func Signout(context *gin.Context) {
	username, ok := ExtractUsernameFromAccessToken(context)
	if !ok {
		return
	}

	err := services.DeleteCachedUserInfo(username)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to sign out"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Signed out successfully"})
}
