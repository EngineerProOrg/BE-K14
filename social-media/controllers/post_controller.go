package controllers

import (
	"net/http"
	"social-media/models"
	"social-media/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreatePost(context *gin.Context) {
	postViewModel := &models.PostViewModel{}

	err := context.ShouldBindJSON(postViewModel)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse request data"})
		return
	}

	userId, ok := ExtractUserIdFromAccessToken(context)
	if !ok {
		return
	}

	postModel := models.CreateMappingPostViewModelToPostEntity(postViewModel)
	postModel.UserId = userId

	createdPost, err := services.CreatePost(postModel)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, gin.H{
		"postId":    createdPost.Id,
		"title":     createdPost.Title,
		"content":   createdPost.Content,
		"createdAt": createdPost.CreatedAt,
		"author":    createdPost.User.Name,
	})
}

func GetPostById(context *gin.Context) {
	postId, err := strconv.ParseInt(context.Param("postId"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "postId": postId})
		return
	}

	postModel, err := services.GetPostById(postId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error(), "postId": postId})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"postId":    postModel.Id,
		"title":     postModel.Title,
		"content":   postModel.Content,
		"createdAt": postModel.CreatedAt,
		"author":    postModel.User.Name,
	})
}
