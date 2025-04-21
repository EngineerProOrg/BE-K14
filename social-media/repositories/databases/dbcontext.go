package databases

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"social-media/utils"

	_ "github.com/go-sql-driver/mysql"
)

var SocialMediaDbContext *sql.DB

func InitSocialMediaDbContext() {
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

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("❌ Cannot open MySQL connection: %v", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatalf("❌ Cannot ping MySQL: %v", err)
	}

	log.Println("✅ Connected to MySQL successfully!")
	SocialMediaDbContext = db

	createTables()
	seedSampleUserData()
}

func createTables() {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS Users 
	(
		Id INT 			AUTO_INCREMENT PRIMARY KEY,
		Name 				VARCHAR(100) NOT NULL,
		Email 			VARCHAR(255) NOT NULL UNIQUE,
		Password 		VARCHAR(255) NOT NULL,
		CreatedAt 	DATETIME DEFAULT CURRENT_TIMESTAMP,
		UpdatedAt 	DATETIME NULL
	);`

	_, err := SocialMediaDbContext.Exec(createUsersTable)
	if err != nil {
		panic(fmt.Sprintf("Could not create table Users. Error: %v", err))
	}
}

func seedSampleUserData() {
	var countUsers int

	query := `SELECT COUNT(*) FROM Users`

	err := SocialMediaDbContext.QueryRow(query).Scan(&countUsers)
	if err != nil {
		fmt.Printf("❌ Could not seed sample data. Error: %v\n", err)
		return
	}

	if countUsers > 0 {
		fmt.Println("✅ Users table already has data. Skipping seeding.")
		return
	}

	exampleHashPassword, _ := utils.HashPassword("P@ssword123")

	insertQuery := `
		INSERT INTO Users(Name, Email, Password, CreatedAt, UpdatedAt)
		VALUES
			('User 001 - Global InfoTrack','user001@infotrack.com.au', ?, NOW(), NULL),
			('User 002 - Global InfoTrack','user002@infotrack.com.au', ?, NOW(), NULL),
			('User 003 - Global InfoTrack','user003@infotrack.com.au', ?, NOW(), NULL),
			('User 004 - Global InfoTrack','user004@infotrack.com.au', ?, NOW(), NULL),
			('User 005 - Global InfoTrack','user005@infotrack.com.au', ?, NOW(), NULL)`

	stmt, err := SocialMediaDbContext.Prepare(insertQuery)
	if err != nil {
		fmt.Printf("❌ Could not prepare statement. Error: %v\n", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		exampleHashPassword,
		exampleHashPassword,
		exampleHashPassword,
		exampleHashPassword,
		exampleHashPassword,
	)

	if err != nil {
		fmt.Printf("❌ Could not execute insert. Error: %v\n", err)
		return
	}
	fmt.Println("✅ Successfully seeded sample data")
}
