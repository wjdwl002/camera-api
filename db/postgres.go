package db

import (
	"fmt"
	"log"

	"camera-app/backend/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Init() *gorm.DB {

	// Log connection details for debugging
	fmt.Printf("Attempting to connect to database:\nHost: %s\nUser: %s\nDatabase: %s\n",
		config.GetHost(),
		config.GetUsername(),
		config.GetDatabase())

	// Construct connection string from individual parameters
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require",
		config.GetHost(),
		config.GetUsername(),
		config.GetPassword(),
		config.GetDatabase(),
		config.GetDBPort(),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}

	DB = db
	fmt.Println("Connected to PostgreSQL via GORM")

	return db
}
