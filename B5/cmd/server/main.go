package main

import (
	"B5/internal/handlers"
	"B5/internal/store"
	"github.com/gin-gonic/gin"
)

func main() {
	redisStore := store.NewRedisStore()
	loginHandler := handlers.NewLoginHandler(redisStore)
	router := gin.Default()
	router.POST("/login", loginHandler.Handle)
	err := router.Run(":8080")
	if err != nil {
		print(err)
	}
}
