package middlewares

import (
	"net/http"
	"social-media/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	accessToken := context.Request.Header.Get("Authorization")

	if len(accessToken) == 0 {
		// something happens then we can stop and no other code on server runs
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// If Authorization contains: "Bearer <token>"
	if strings.HasPrefix(accessToken, "Bearer ") {
		accessToken = strings.TrimPrefix(accessToken, "Bearer ")
	}

	userId, username, err := utils.VerifyToken(accessToken)
	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}
	// set userid into gin.context
	context.Set("userId", userId)
	// set username into gin.context
	context.Set("username", username)

	// this ensures that next request in line will execute correctly
	context.Next()
}
