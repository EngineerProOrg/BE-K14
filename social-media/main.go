package main

import (
	"log"
	"social-media/repositories/databases"
	"social-media/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	setupMySqlDatabase()
	setupRedisClient()
	registerRoutes()
}

func setupMySqlDatabase() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Error loading .env file")
		return
	}
	// Init MySql
	databases.InitGormContext()
}

func setupRedisClient() {
	databases.InitRedisClient()
}

func registerRoutes() {
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080") //localhost:8080
}
