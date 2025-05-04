package controllers

import (
	"errors"
	"net/http"
	"social-media/models"
	"social-media/services"
	"social-media/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Signup(context *gin.Context) {
	// Validate and bind request body
	userSignupViewModel, ok := utils.BindAndValidate[models.UserSignupRequestViewModel](context)
	if !ok {
		return
	}

	user := models.MapUserSignupRequestViewModelToUserDbModel(userSignupViewModel)

	err := services.Signup(user)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"userId": user.ID})
}

func Signin(context *gin.Context) {
	userInput, ok := utils.BindAndValidate[models.UserSigninRequestViewModel](context)
	if !ok {
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

	// Cached user_info after login success
	// Therefore, we can save time to querydb and don't need to call Preload("User")
	services.SetCachedUserInfo(context, userSigninResponseVm.UserId, userSigninResponseVm)
	context.JSON(http.StatusOK, gin.H{"message": "success", "access_token": accessToken})
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
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: userId mismatch"})
		return
	}

	userSigninResponseVm, err := services.GetCachedUserInfo(context, userId)
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
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized: userId mismatch"})
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
	services.SetCachedUserInfo(context, userId, updatedProfile)
	context.JSON(http.StatusOK, gin.H{"userInfo": updatedProfile})
}

func Signout(context *gin.Context) {
	userId, ok := ExtractUserIdFromAccessToken(context)
	if !ok {
		return
	}

	err := services.DeleteCachedUserInfo(context, userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to sign out"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Signed out successfully"})
}
