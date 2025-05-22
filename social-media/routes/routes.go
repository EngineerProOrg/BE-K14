package routes

import (
	"social-media/controllers"
	"social-media/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(engine *gin.Engine) {
	apiGroup := engine.Group("/api/v1")

	registerPublicRoutes(apiGroup)
	registerProtectedRoutes(apiGroup)
}

// Public routes without authen
func registerPublicRoutes(router *gin.RouterGroup) {
	router.POST("/users/signup", controllers.Signup)
	router.POST("/users/signin", controllers.Signin)
}

// Protected routes (must Auth)
func registerProtectedRoutes(router *gin.RouterGroup) {
	protected := router.Group("")
	protected.Use(middlewares.Authenticate)

	// User endpoints
	protected.GET("/users/profile/:userId", controllers.GetUserProfile)
	protected.PUT("/users/profile/:userId", controllers.EditUserProfile)
	protected.GET("/users/:userId/posts", controllers.GetPostsByUserId)

	// Post endpoints
	protected.POST("/posts", controllers.CreatePost)
	protected.PUT("/posts/:postId", controllers.UpdatePost)
	protected.GET("/posts/:postId", controllers.GetPostById)
	protected.GET("/posts", controllers.GetPosts)

	// Comment endpoints
	protected.POST("/posts/:postId/comments", controllers.CreateComment)
	protected.PUT("/posts/:postId/comments/:commentId", controllers.UpdateComment)
	protected.GET("/posts/:postId/comments", controllers.GetCommentsByPostId)

	// Reaction endpoints
	protected.GET("/reactions/:targetId", controllers.GetReactionsByTarget) // use query param: ?target_type=post|comment
	protected.POST("/reactions", controllers.CreateOrUpdateReaction)
	protected.GET("/reactions/:targetId/grouped", controllers.CountGroupedReactionsByTarget) // use query param: ?target_type=post|comment

	protected.POST("follow/friends/:userId", controllers.FollowUser)
	protected.DELETE("follow/friends/:userId", controllers.UnfollowUser)
	protected.GET("follow/friends/:userId", controllers.GetFollowings)
}
