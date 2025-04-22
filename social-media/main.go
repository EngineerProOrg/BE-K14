package main

import (
	"log"
	"social-media/repositories/databases"

	"github.com/joho/godotenv"
)

func main() {
	setupMySqlDatabase()
	setupRedisClient()
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
