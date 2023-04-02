package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	Models "go_practice/models"
	Router "go_practice/routers"
)

func main() {
	fmt.Println("=========API Started=========")

	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get database configuration values from environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Connect to the database
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	// Auto migrate User and Login tables
	user := Models.User{}
	login := Models.Login{}
	err = db.AutoMigrate(&user, &login)
	if err != nil {
		log.Fatal("failed to migrate tables")
	}

	// Start the API server
	router := Router.Router(db)
	router.Run(":9090")
}
