package main

import (
	"fmt"

	Router "go_practice/routers"

	Models "go_practice/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	fmt.Println("=========API Started=========")

	// Connect to the database
	dsn := "host=localhost user=postgres password=postgres dbname=golang_practice port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// create the database
	// Create the database if it does not exist
	db.Exec("CREATE DATABASE  golang_practice;")

	// // Use the golang_practice database
	// dsn = "host=localhost user=postgres password=postgres dbname=golang_practice port=5432 sslmode=disable"
	// sqlDB, err := db.DB()
	// if err != nil {
	// 	panic("failed to get sql.DB")
	// }
	// sqlDB.Close()

	// Auto migrate User and Login tables
	user := Models.User{}
	login := Models.Login{}
	err = db.AutoMigrate(&user, &login)
	if err != nil {
		panic("failed to migrate tables")
	}

	// Start the API server
	router := Router.Router(db)
	router.Run("localhost:9090")
}
