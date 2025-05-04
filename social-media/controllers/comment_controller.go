package controllers

import (
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

	postId, err := strconv.ParseInt(c.Param("postId"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	commentRequestViewModel, ok := utils.BindAndValidate[models.CommentRequestViewModel](c)
	if !ok {
		return
	}

	comment := &models.Comment{
		Content: commentRequestViewModel.Content,
		UserId:  userId,
		PostId:  postId,
	}

	createdComment, err := services.CreateComment(comment)
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
