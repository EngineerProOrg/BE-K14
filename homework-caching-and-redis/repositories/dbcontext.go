package repositories

import (
	"database/sql"
	"fmt"
	"homework-caching-and-redis/utils"

	_ "github.com/mattn/go-sqlite3"
)

var DbContext *sql.DB

func InitDatabaseContext() {
	var err error

	DbContext, err = sql.Open("sqlite3", "repositories/database/HomeworkCachingAndRedis.db")

	if err != nil {
		panic("Could not establish connection to database")
	}

	DbContext.SetMaxOpenConns(10)
	DbContext.SetMaxIdleConns(5)

	createTables()
	seedSampleUserData()
}

func createTables() {
	createUsersTable := `
		CREATE TABLE IF NOT EXISTS users(
			id						INTEGER PRIMARY KEY AUTOINCREMENT,
			name					TEXT NOT NULL,
			email					TEXT NOT NULL UNIQUE,
			password			TEXT NOT NULL,
			createdAt			TEXT NOT NULL,
			updatedAt			TEXT)`

	_, err := DbContext.Exec(createUsersTable)
	if err != nil {
		panic(fmt.Sprintf("Could not create table users. Error: %v", err))
	}
}

func seedSampleUserData() {
	var countUsers int

	query := `SELECT COUNT (*) FROM users`

	err := DbContext.QueryRow(query).Scan(&countUsers)
	if err != nil {
		fmt.Printf("❌ Could not seed sample data. Error: %v", err)
		return
	}

	if countUsers > 0 {
		fmt.Println("✅ users table already has data. Skipping seeding.")
		return
	}

	exampleHashPassword, _ := utils.HashPassword("P@ssword123")

	insertQuery := `
		INSERT INTO users(name, email, password, createdAt, updatedAt)
		VALUES 
			('User 001 - Global InfoTrack','user001@infotrack.com.au', ?, datetime('now'), null),
			('User 002 - Global InfoTrack','user002@infotrack.com.au', ?, datetime('now'), null),
			('User 003 - Global InfoTrack','user003@infotrack.com.au', ?, datetime('now'), null),
			('User 004 - Global InfoTrack','user004@infotrack.com.au', ?, datetime('now'), null),
			('User 005 - Global InfoTrack','user005@infotrack.com.au', ?, datetime('now'), null)`
	stmt, err := DbContext.Prepare(insertQuery)
	if err != nil {
		fmt.Printf("❌ Could not seed sample data. Error: %v", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(exampleHashPassword,
		exampleHashPassword,
		exampleHashPassword,
		exampleHashPassword,
		exampleHashPassword)

	if err != nil {
		fmt.Printf("❌ Could not execute query. Error: %v", err)
		return
	}
	fmt.Println("✅ Successfully seeded sample data")
}
