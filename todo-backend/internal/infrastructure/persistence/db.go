package persistence

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// ConnectDB opens a connection to PostgreSQL using the connection string from the DB_DSN environment variable.
// Returns an instance of *gorm.DB or error if connection failed.
func ConnectDB() (*gorm.DB, error) {
	dsn := os.Getenv("DB_DSN")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}
