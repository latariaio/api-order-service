package database

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using system env vars")
	}

	// postgresURI := os.Getenv("DATABASE_URL")
	dsn := "host=postgres user=postgres password=postgres dbname=app port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Panic(err)
	}

	if err := sqlDB.Ping(); err != nil {
		log.Panic(err)
	}

	fmt.Println("Connected to database")

	return db
}
