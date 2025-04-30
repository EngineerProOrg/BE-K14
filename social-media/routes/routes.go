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
	protected.GET("/posts/:postId", controllers.GetPostById)
	protected.GET("/posts", controllers.GetPosts)

	// Comment endpoints
	protected.POST("/posts/:postId/comments", controllers.CreateComment)
	protected.GET("/posts/:postId/comments", controllers.GetCommentsByPostId)

}
