package middlewares

import (
	"net/http"
	"social-media/utils"

	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	accessToken := context.Request.Header.Get("Authorization")

	if len(accessToken) == 0 {
		// something happens then we can stop and no other code on server runs
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userId, err := utils.VerifyToken(accessToken)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	context.Set("userId", userId)

	// this ensures that next request in line will execute correctly
	context.Next()
}
