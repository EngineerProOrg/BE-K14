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

func registerPublicRoutes(router *gin.RouterGroup) {
	// Public routes without authen
	router.POST("/users/signup", controllers.Signup)
	router.POST("/users/signin", controllers.Signin)
}

func registerProtectedRoutes(router *gin.RouterGroup) {
	// Protected routes (must Auth)
	protected := router.Group("")
	protected.Use(middlewares.Authenticate)

	protected.GET("/users/profile/:userId", controllers.GetUserProfile)
	protected.PUT("/users/profile/:userId", controllers.EditUserProfile)
	protected.GET("/users/:userId/posts", controllers.GetPostsByUserId)

	protected.POST("/posts", controllers.CreatePost)
	protected.GET("/posts/:postId", controllers.GetPostById)
	protected.GET("/posts", controllers.GetPosts)
}
