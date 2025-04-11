package main

import (
	"B5/internal/handlers"
	"B5/internal/middleware"
	"B5/internal/store"
	"github.com/gin-gonic/gin"
)

func main() {
	redisStore := store.NewRedisStore()
	loginHandler := handlers.NewLoginHandler(redisStore)
	pingHandler := handlers.NewPingHandler(redisStore)
	topHandler := handlers.NewTopHandler(redisStore)
	countHandler := handlers.NewCountHandler(redisStore)

	sessionMiddleware := middleware.NewSessionMiddleware(redisStore)

	router := gin.Default()
	router.POST("/login", loginHandler.Handle)

	protected := router.Group("/")
	protected.Use(sessionMiddleware.ValidateSession())
	{
		protected.POST("/ping", pingHandler.Handle)
		protected.GET("/top", topHandler.Handle)
		protected.GET("/count", countHandler.Handle)
	}

	err := router.Run(":8080")
	if err != nil {
		print(err)
	}
}
