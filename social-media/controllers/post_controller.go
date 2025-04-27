package controllers

import (
	"net/http"
	"social-media/models"
	"social-media/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreatePost(context *gin.Context) {
	postViewModel := &models.PostRequestViewModel{}

	err := context.ShouldBindJSON(postViewModel)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse request data"})
		return
	}

	userId, ok := ExtractUserIdFromAccessToken(context)
	if !ok {
		return
	}

	postModel := models.CreateMappingPostRequestViewModelToPostEntity(postViewModel)
	postModel.UserId = userId

	responsePostvm, err := services.CreatePost(postModel)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"post": responsePostvm})
}

func GetPostById(context *gin.Context) {
	postId, err := strconv.ParseInt(context.Param("postId"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "postId": postId})
		return
	}

	postResponse, err := services.GetPostById(postId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error(), "postId": postId})
		return
	}
	context.JSON(http.StatusOK, gin.H{"post": postResponse})
}

func GetPosts(context *gin.Context) {
	posts, err := services.GetPosts()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"posts": posts})
}
