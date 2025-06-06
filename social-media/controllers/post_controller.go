package controllers

import (
	"net/http"
	"social-media/models"
	"social-media/services"
	"social-media/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreatePost(context *gin.Context) {
	// Validate and bind request body
	postRequestViewModel, ok := utils.BindAndValidate[models.PostRequestViewModel](context)
	if !ok {
		return
	}

	// Extract userId from token
	userId, ok := ExtractUserIdFromAccessToken(context)
	if !ok {
		return
	}

	username, ok := ExtractUsernameFromAccessToken(context)
	if !ok {
		return
	}

	postRequestViewModel.UserId = userId
	postRequestViewModel.Username = username

	// call service to create post
	responsePostvm, err := services.CreatePost(postRequestViewModel)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// return result
	context.JSON(http.StatusCreated, responsePostvm)
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
	context.JSON(http.StatusOK, postResponse)
}

func GetPosts(context *gin.Context) {
	posts, err := services.GetPosts()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, posts)
}

func GetPostsByUserId(context *gin.Context) {
	userId, err := strconv.ParseInt(context.Param("userId"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "userId": userId})
		return
	}

	extractedUserId, ok := ExtractUserIdFromAccessToken(context)
	if !ok {
		return
	}
	if extractedUserId != userId {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	extractedUsername, ok := ExtractUsernameFromAccessToken(context)
	if !ok {
		return
	}

	postResponseVm, err := services.GetPostsByUserId(userId, extractedUsername)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, postResponseVm)
}

func UpdatePost(c *gin.Context) {
	postIdParam := c.Param("postId")
	postId, err := strconv.ParseInt(postIdParam, 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid postId"})
		return
	}
	userId, ok := ExtractUserIdFromAccessToken(c)
	if !ok {
		return
	}

	updateReq, ok := utils.BindAndValidate[models.PostRequestViewModel](c)
	if !ok {
		return
	}

	// check post exist
	postResponseVm, err := services.GetPostById(postId)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error(), "postId": postId})
		return
	}

	if postResponseVm.Author.UserId != userId {
		c.JSON(http.StatusForbidden, gin.H{"error": "You do not have permission to update this post"})
		return
	}

	err = services.UpdatePost(postId, userId, updateReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "updated success"})
}
