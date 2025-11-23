package config

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
)

var DB *gorm.DB

func InitDataBase() (*gorm.DB, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		dbPort = "5433"
	}
	dbSslMode := os.Getenv("DB_SSLMODE")

	connStr := "host=localhost user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " port=" + dbPort + " sslmode=" + dbSslMode

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	return db, nil

}
