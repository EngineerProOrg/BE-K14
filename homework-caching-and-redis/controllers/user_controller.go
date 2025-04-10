package controllers

import (
	"homework-caching-and-redis/models"
	"homework-caching-and-redis/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignUp(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse request data"})
	}

	err = services.CreateANewUser(&user)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	context.JSON(http.StatusCreated, gin.H{"userId": user.Id})
}

func LoginIn(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse request data"})
		return
	}

	err = services.Login(&user)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	sessionId, err := services.GenerateSessionId(context, user.Email)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not store session in Redis"})
		return
	}

	context.SetCookie("session_id", sessionId, 3600, "/", "", false, true)
	context.JSON(http.StatusOK, gin.H{"message": "Login success"})
}
