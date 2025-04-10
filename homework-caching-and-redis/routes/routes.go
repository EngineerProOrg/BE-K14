package routes

import (
	"homework-caching-and-redis/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(ginEngine *gin.Engine) {
	registerUserRoute(ginEngine)
}

func registerUserRoute(ginEngine *gin.Engine) {
	ginEngine.POST("/signup", controllers.SignUp)
	ginEngine.POST("/login", controllers.LoginIn)
}
