package routes

import (
	"social-media/controllers"
	"social-media/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(engine *gin.Engine) {
	registerUserRoutes(engine)
}

func registerUserRoutes(ginEngine *gin.Engine) {
	userGroup := ginEngine.Group("api/v1")
	userGroup.POST("/users/signup", controllers.Signup)
	userGroup.POST("/users/signin", controllers.Signin)

	// Call middleware
	userGroup.Use(middlewares.Authenticate)
	userGroup.GET("/users/profile/:userId", controllers.GetUserProfile)
	userGroup.PUT("/users/profile/:userId", controllers.EditUserProfile)
}
