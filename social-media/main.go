package main

import (
	"log"
	"social-media/repositories/databases"
	"social-media/routes"
	"social-media/utils"
	"time"

	"github.com/gin-contrib/cors"
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
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"}, // FE port
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	routes.RegisterRoutes(server)
	server.Run(":8080") // BE port
}

func registerCustomValidation() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("notblank", utils.NotBlank)
	}
}
