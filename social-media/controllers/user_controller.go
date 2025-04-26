package controllers

import (
	"net/http"
	"social-media/models"
	"social-media/services"
	"social-media/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Signup(context *gin.Context) {
	// notes var userSignupViewModel *models.UserSignupViewModel // khai báo con trỏ nhưng chưa gán địa chỉ

	userSignupViewModel := &models.UserSignupViewModel{} // best practice -> Tạo struct mới rồi lấy địa chỉ luôn ->Đã trỏ tới 1 struct rỗng

	err := context.ShouldBindJSON(userSignupViewModel)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse request data"})
		return
	}

	user := models.CreateMappingUserSignupViewModelToUserEntity(userSignupViewModel)

	err = services.Signup(user)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"userId": user.ID})
}

func Signin(context *gin.Context) {
	userInput := &models.User{}

	err := context.ShouldBindJSON(userInput)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse request data"})
		return
	}
	err = services.Signin(userInput)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	accessToken, err := utils.GenerateAccessToken(userInput.Email, userInput.ID)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "success", "access_token": accessToken})
}

func GetUserProfile(context *gin.Context) {
	userId, err := strconv.ParseInt(context.Param("userId"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "userId": userId})
		return
	}

	user, err := services.GetUserProfile(userId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error(), "userId": userId})
		return
	}

	userProfile := user.CreateMapingUserEntityToCreateProfileViewModel()
	context.JSON(http.StatusOK, gin.H{"userProfile": userProfile})
}

func EditUserProfile(context *gin.Context) {
	userId, err := strconv.ParseInt(context.Param("userId"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "userId": userId})
		return
	}

	_, err = services.GetUserProfile(userId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error(), "userId": userId})
		return
	}

	editUserProfile := &models.EditUserProfileViewModel{}
	err = context.ShouldBindJSON(editUserProfile)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse request data"})
		return
	}

	err = services.EditUserProfile(userId, editUserProfile)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "userId": userId})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "success"})
}
