package controllers

import (
	"net/http"
	"social-media/models"
	"social-media/services"

	"github.com/gin-gonic/gin"
)

func CreatePost(context *gin.Context) {
	postViewModel := &models.PostViewModel{}

	err := context.ShouldBindJSON(postViewModel)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse request data"})
		return
	}

	postModel := models.CreateMappingPostViewModelToPostEntity(postViewModel)

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
