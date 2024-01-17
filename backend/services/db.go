package services

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var dbConnect *gorm.DB

// InitDBConnection initializes the database connection.
//
// It retrieves the necessary environment variables for the database connection: DB_USER, DB_PASS, DB_NAME, DB_HOST, and DB_PORT.
// Then it creates a DSN (Data Source Name) string using the retrieved environment variables.
// Next, it opens a connection to the database using the gorm package and the created DSN string.
// If there is an error during the connection process, it panics with the message "failed to connect to database".
func InitDBConnection() {
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Minsk", dbHost, dbUser, dbPass, dbName, dbPort)

	var err error
	dbConnect, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
}

// GetDBConnection returns the database connection.
//
// No parameters.
// Returns a pointer to a gorm.DB object.
func GetDBConnection() *gorm.DB {
	return dbConnect
}
