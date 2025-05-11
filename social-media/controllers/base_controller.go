package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ExtractUserIdFromAccessToken(context *gin.Context) (int64, bool) {
	userIdAny, exists := context.Get("userId")
	if !exists {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "User not authorized"})
		return 0, false
	}

	userId, ok := userIdAny.(int64)
	if !ok {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid userId type"})
		return 0, false
	}

	return userId, true
}

func ExtractUsernameFromAccessToken(context *gin.Context) (string, bool) {
	usernameAny, exists := context.Get("username")
	if !exists {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "User not authorized"})
		return "", false
	}

	userId, ok := usernameAny.(string)
	if !ok {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username type"})
		return "", false
	}

	return userId, true
}
