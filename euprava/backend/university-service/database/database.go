package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type Env struct {
	DB *sql.DB
}

func ConnectDB() (*sql.DB, error) {
	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env")
	}

	connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return nil, err
	}

	return db, nil
}
