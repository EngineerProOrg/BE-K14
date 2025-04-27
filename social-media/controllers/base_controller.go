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
