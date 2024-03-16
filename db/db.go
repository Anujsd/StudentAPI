package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	connStr := os.Getenv("CONNECTION_STRING")

	// Connect to the database
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	// defer DB.Close()

	// Test the connection
	err = DB.Ping()
	if err != nil {
		log.Fatal("Error pinging the database:", err)
	}

	fmt.Println("Connected to the database")

	// Perform database operations here
	createTables()
}

func createTables() {
	createStudentsTable := `
		CREATE TABLE IF NOT EXISTS students (
			ID SERIAL PRIMARY KEY,
			FirstName VARCHAR(255),
			LastName VARCHAR(255),
			EmailId VARCHAR(255)
		)
	`
	_, err := DB.Exec(createStudentsTable)

	if err != nil {
		fmt.Print(err)
		panic("Could not create students table.")
	}
}
