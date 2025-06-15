package main

import (
	"homework-caching-and-redis/repositories"
	"homework-caching-and-redis/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	repositories.InitDatabaseContext()
	repositories.InitRedisClient()

	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080") //localhost:8080
}
