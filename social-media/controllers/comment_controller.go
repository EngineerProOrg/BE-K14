package controllers

import (
	"errors"
	"net/http"
	"social-media/models"
	"social-media/services"
	"social-media/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateComment(c *gin.Context) {
	userId, ok := ExtractUserIdFromAccessToken(c)
	if !ok {
		return
	}
	username, ok := ExtractUsernameFromAccessToken(c)
	if !ok {
		return
	}

	postId, err := strconv.ParseInt(c.Param("postId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	commentRequestViewModel, ok := utils.BindAndValidate[models.CommentRequestViewModel](c)
	if !ok {
		return
	}
	commentRequestViewModel.UserId = userId
	commentRequestViewModel.Username = username
	commentRequestViewModel.PostId = postId

	createdComment, err := services.CreateComment(commentRequestViewModel)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"comment": createdComment})
}

func GetCommentsByPostId(c *gin.Context) {
	postId, err := strconv.ParseInt(c.Param("postId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	comments, err := services.GetCommentsByPostId(postId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"comments": comments})
}

func UpdateComment(c *gin.Context) {
	userId, ok := ExtractUserIdFromAccessToken(c)
	if !ok {
		return
	}
	username, ok := ExtractUsernameFromAccessToken(c)
	if !ok {
		return
	}

	postId, err := strconv.ParseInt(c.Param("postId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}
	commentId, err := strconv.ParseInt(c.Param("commentId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid comment ID"})
		return
	}

	commentRequestViewModel, ok := utils.BindAndValidate[models.CommentRequestViewModel](c)
	if !ok {
		return
	}

	commentRequestViewModel.PostId = postId
	commentRequestViewModel.UserId = userId
	commentRequestViewModel.CommentId = commentId
	commentRequestViewModel.Username = username

	commentResponseVm, err := services.UpdateComment(commentRequestViewModel)
	if errors.Is(err, utils.ErrCannotEditComment) {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	if errors.Is(err, utils.ErrCommentNotInPost) {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"comment": commentResponseVm})
}
