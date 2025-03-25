package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func Database() *sql.DB {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Retrieve database credentials from environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")

	// Construct database connection string
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/?parseTime=true",
		dbUser,
		dbPass,
		dbHost,
		dbPort,
	)

	// Open database connection
	database, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
	}

	// Verify database connection
	if err = database.Ping(); err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// Create database if not exists
	_, err = database.Exec(`CREATE DATABASE IF NOT EXISTS gotodo`)
	if err != nil {
		log.Fatalf("Error creating database: %v", err)
	}

	// Use the database
	_, err = database.Exec(`USE gotodo`)
	if err != nil {
		log.Fatalf("Error using database: %v", err)
	}

	// Create table if not exists
	_, err = database.Exec(`
		CREATE TABLE IF NOT EXISTS todos (
		    id INT AUTO_INCREMENT,
		    item TEXT NOT NULL,
		    completed BOOLEAN DEFAULT FALSE,
		    PRIMARY KEY (id)
		);
	`)
	if err != nil {
		log.Fatalf("Error creating table: %v", err)
	}

	fmt.Println("Database Connection Successful")
	return database
}
