package controllers

import (
	"net/http"
	"social-media/models"
	"social-media/services"
	"social-media/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetReactionsByTarget(context *gin.Context) {
	targetId, err := strconv.ParseInt(context.Param("targetId"), 10, 64)
	targetType := context.Query("target_type")

	if err != nil || (targetType != "post" && targetType != "comment") {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid target"})
		return
	}

	reactions, err := services.GetReactionsByTarget(targetId, targetType)
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

	reactionRequestVm, ok := utils.BindAndValidate[models.ReactionRequestViewModel](context)
	if !ok {
		return
	}

	reactionRequestVm.UserId = extractedUserId

	err := services.CreateOrUpdateReaction(reactionRequestVm)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "success"})
}

func CountGroupedReactionsByTarget(context *gin.Context) {
	_, ok := ExtractUserIdFromAccessToken(context)
	if !ok {
		return
	}
	targetId, err := strconv.ParseInt(context.Param("targetId"), 10, 64)
	targetType := context.Query("target_type")
	if err != nil || (targetType != "post" && targetType != "comment") {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid target"})
		return
	}

	reactionCount, err := services.CountGroupedReactionsByTarget(targetId, targetType)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, reactionCount)
}
