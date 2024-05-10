package data

import (
	"database/sql"
	"fmt"
	"log"
)

type DormRepo struct {
	logger *log.Logger
	DB     *sql.DB
}

// database variables
// usually we should get them from env like os.Getenv("variableName")
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "root"
	dbname   = "dorm"
)

// ConnectDB tries to connect DB and on succcesful it returns
// DB connection string and nil error, otherwise return empty DB and the corresponding error.
func ConnectDB() (*sql.DB, error) {

	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname = %s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", connString)
	if err != nil {
		log.Printf("failed to connect to database: %v", err)
		return &sql.DB{}, err
	}
	return db, nil
}
