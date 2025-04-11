package main

import (
	"B5/internal/handlers"
	"B5/internal/store"
	"github.com/gin-gonic/gin"
)

func main() {
	redisStore := store.NewRedisStore()
	loginHandler := handlers.NewLoginHandler(redisStore)
	pingHandler := handlers.NewPingHandler(redisStore)
	topHandler := handlers.NewTopHandler(redisStore)
	countHandler := handlers.NewCountHandler(redisStore)
	router := gin.Default()
	router.POST("/login", loginHandler.Handle)
	router.POST("/ping", pingHandler.Handle)
	router.GET("/top", topHandler.Handle)
	router.GET("/count", countHandler.Handle)
	err := router.Run(":8080")
	if err != nil {
		print(err)
	}
}
