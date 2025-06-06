package controllers

import (
	"net/http"
	"social-media/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FollowUser(c *gin.Context) {
	var followService = services.NewFollowService()

	followerID, ok := ExtractUserIdFromAccessToken(c)
	if !ok {
		return
	}

	targetIDStr := c.Param("userId")
	targetID, err := strconv.ParseInt(targetIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	err = followService.Follow(followerID, targetID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Followed successfully"})
}

func UnfollowUser(c *gin.Context) {
	var followService = services.NewFollowService()

	followerID, ok := ExtractUserIdFromAccessToken(c)
	if !ok {
		return
	}
	targetIDStr := c.Param("userId")
	targetID, err := strconv.ParseInt(targetIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	err = followService.Unfollow(followerID, targetID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Unfollowed successfully"})
}

func GetFollowings(c *gin.Context) {
	var followService = services.NewFollowService()

	userIDStr := c.Param("userId")
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	followings, err := followService.GetFollowings(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"followings": followings})
}
