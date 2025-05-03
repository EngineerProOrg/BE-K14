package databases

import (
	"fmt"
	"log"
	"os"
	"social-media/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var GormDb *gorm.DB

func InitGormContext() {
	user := os.Getenv("MYSQL_USER")
	pass := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	dbname := os.Getenv("MYSQL_DB")

	if user == "" || pass == "" || host == "" || port == "" || dbname == "" {
		log.Fatal("❌ Missing MySQL environment variables")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		user, pass, host, port, dbname)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Failed to connect to MySQL via GORM: %v", err)
	}
	log.Println("✅ Connected to MySQL via GORM!")

	// Auto-migrate models
	err = db.AutoMigrate(
		&models.User{},
		&models.Post{},
		&models.Reaction{},
		&models.Comment{},
		&models.Follow{},
	)
	if err != nil {
		log.Fatalf("❌ Auto migration failed: %v", err)
	}

	GormDb = db

	//ClearAllData(db)

	seedSampleUserData(db)
	seedSamplePostData(db)
	seedSampleCommentData(db)
	seedSampleReactionData(db)
}
