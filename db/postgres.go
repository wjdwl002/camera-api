package db

import (
	"database/sql"
	"log"

	"camera-app/backend/config"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Init() {
	dsn := config.GetDSN()

	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatal("Failed to ping DB:", err)
	}

	log.Println("Connected to PostgreSQL")
}
