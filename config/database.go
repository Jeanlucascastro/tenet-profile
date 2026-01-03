package config

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/joho/godotenv"

	"tenet-profile/internal/model"
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
	dbHost := os.Getenv("DB_HOST")
	if dbPort == "" {
		dbPort = "5433"
	}
	dbSslMode := os.Getenv("DB_SSLMODE")

	connStr := "host=" + dbHost + " user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " port=" + dbPort + " sslmode=" + dbSslMode

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	return db, nil

}

func RunMigrations(db *gorm.DB) {

	if db == nil {
		log.Fatal("Database connection is nil")
	}

	for _, model := range model.Models {
		err := db.AutoMigrate(model)
		if err != nil {
			log.Fatalf("Failed to migrate model %T: %v", model, err)
		}
	}

	log.Println("Database migrations completed successfully.")
}
