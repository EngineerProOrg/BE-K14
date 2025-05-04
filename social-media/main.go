package main

import (
	"log"
	"social-media/repositories/databases"
	"social-media/routes"
	"social-media/utils"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

func main() {
	setupMySqlDatabase()
	setupRedisClient()
	registerCustomValidation()
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

func registerCustomValidation() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("notblank", utils.NotBlank)
	}
}
