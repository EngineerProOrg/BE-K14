package routes

import (
	"social-media/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(engine *gin.Engine) {
	registerUserRoutes(engine)
}

func registerUserRoutes(ginEngine *gin.Engine) {
	userGroup := ginEngine.Group("api/v1")
	{
		userGroup.POST("/users", controllers.Signup)
	}
}
