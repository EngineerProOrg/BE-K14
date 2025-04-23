package controllers

import (
	"net/http"
	"social-media/models"
	"social-media/services"

	"github.com/gin-gonic/gin"
)

func Signup(context *gin.Context) {
	var userSignupViewModel models.UserSignupViewModel

	err := context.ShouldBindJSON(&userSignupViewModel)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse request data"})
		return
	}

	user, err := models.MapToUserEntity(userSignupViewModel)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	err = services.Signup(&user)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"userId": user.ID})
}

func Signin(context *gin.Context) {
	var userInput models.User

	err := context.ShouldBindJSON(&userInput)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse request data"})
		return
	}
	err = services.Signin(&userInput)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "success"})
}
