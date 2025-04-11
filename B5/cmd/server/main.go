package main

import (
	"B5/internal/config"
	"B5/internal/handlers"
	"B5/internal/middleware"
	"B5/internal/store"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.NewConfig()
	redisStore := store.NewRedisStore(cfg.RedisAddr, cfg.RedisPassword, cfg.RedisDB)
	sessionMiddleware := middleware.NewSessionMiddleware(redisStore)
	rateLimiter := middleware.NewRateLimiter(redisStore, cfg.RateLimit.Requests, cfg.RateLimit.Window)

	loginHandler := handlers.NewLoginHandler(redisStore)
	pingHandler := handlers.NewPingHandler(redisStore)
	topHandler := handlers.NewTopHandler(redisStore)
	countHandler := handlers.NewCountHandler(redisStore)

	router := gin.Default()
	router.POST("/login", loginHandler.Handle)

	protected := router.Group("/")
	protected.Use(sessionMiddleware.ValidateSession())
	{
		protected.POST("/ping", rateLimiter.Limit(), pingHandler.Handle)
		protected.GET("/top", topHandler.Handle)
		protected.GET("/count", countHandler.Handle)
	}

	err := router.Run(cfg.ServerPort)
	if err != nil {
		print(err)
	}
}
