package controllers

import (
	"net/http"
	"social-media/models"
	"social-media/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetReactionsByPostId(context *gin.Context) {
	postId, err := strconv.ParseInt(context.Param("postId"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid postId"})
		return
	}

	reactions, err := services.GetUserReactionsByPostId(context, postId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, reactions)
}

func CreateOrUpdateReaction(context *gin.Context) {
	extractedUserId, ok := ExtractUserIdFromAccessToken(context)
	if !ok {
		return
	}

	reactionRequestVm := &models.ReactionRequestViewModel{}
	reactionRequestVm.UserId = extractedUserId
	err := context.ShouldBindJSON(reactionRequestVm)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse request data"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "success"})
}
