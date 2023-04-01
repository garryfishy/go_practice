package service

import (
	"fmt"

	"github.com/gin-gonic/gin"

	Helpers "go_practice/helpers"
	Model "go_practice/models"

	"gorm.io/gorm"
)

func InsertUser(c *gin.Context, db *gorm.DB, model Model.User) bool {
	payload := &model
	fmt.Println(payload, "<<<ini palod")

	if result := db.Create(payload); result.Error != nil {
		fmt.Println("Error creating user:", result.Error)
		return false
	}

	return true
}

func InsertLogin(c *gin.Context, db *gorm.DB, model Model.Login) bool {

	payload := &model
	hashed, err := Helpers.HashPassword(payload.Password)
	if err != nil {
		fmt.Println("Error hashing password in InsertLogin:", err)
		return false
	}
	model.Password = hashed

	if result := db.Create(payload); result.Error != nil {
		fmt.Println("Error creating record in InsertLogin:", result.Error)
		return false
	}

	return true
}
