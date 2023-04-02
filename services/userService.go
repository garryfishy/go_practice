package service

import (
	"fmt"

	"github.com/gin-gonic/gin"

	Helpers "go_practice/helpers"
	Model "go_practice/models"

	"gorm.io/gorm"
)

func InsertUser(c *gin.Context, db *gorm.DB, model Model.User) (uint, error) {
	payload := &model
	// validate the user model
	if err := Helpers.ValidateUserModel(payload); err != nil {
		fmt.Println("Error validating user model:", err)
		return 0, err
	}

	if result := db.Create(payload); result.Error != nil {
		fmt.Println("Error creating user:", result.Error)
		return 0, result.Error
	}

	return payload.ID, nil
}

func InsertLogin(c *gin.Context, db *gorm.DB, model Model.Login, userID uint) bool {
	payload := &model
	if err := Helpers.ValidateLogin(payload); err != nil {
		fmt.Println("Errir validating login")
		return false
	}
	hashed, err := Helpers.HashPassword(payload.Password)
	if err != nil {
		fmt.Println("Error hashing password in InsertLogin:", err)
		return false
	}
	model.Password = hashed
	model.UserID = userID

	if result := db.Create(payload); result.Error != nil {
		fmt.Println("Error creating record in InsertLogin:", result.Error)
		return false
	}

	return true
}
