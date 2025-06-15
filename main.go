package main

import (
	"log"
	"os"

	"camera-app/backend/db"
	"camera-app/backend/routes"

	"github.com/joho/godotenv"
)

func main() {
	// Load env vars
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system envs")
	}

	// Init DB
	dbConn := db.Init()

	// Init router
	r := routes.InitRouter(dbConn)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Printf("API running on port %s\n", port)
	r.Run(":" + port)
}
