package service

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	Helpers "go_practice/helpers"
	Model "go_practice/models"
	Structs "go_practice/structs"

	"gorm.io/gorm"
)

func InsertUser(c *gin.Context, db *gorm.DB, model Model.User) (uint, error, Structs.Response) {
	payload := &model
	response := Structs.Response{}
	// validate the user model
	if err := Helpers.ValidateUserModel(payload); err != nil {
		response.Code = http.StatusBadRequest
		response.Message = err.Error()
		return 0, err, response
	}

	if result := db.Create(payload); result.Error != nil {
		response.Code = http.StatusBadRequest
		response.Message = result.Error.Error()
		return 0, result.Error, response
	}

	return payload.ID, nil, response
}

func InsertLogin(c *gin.Context, db *gorm.DB, model Model.Login, userID uint) (bool, Structs.Response) {
	payload := &model
	response := Structs.Response{}
	if err := Helpers.ValidateLogin(payload); err != nil {
		fmt.Println("Errir validating login")
		response.Code = http.StatusBadRequest
		response.Message = err.Error()
		return false, response
	}
	hashed, err := Helpers.HashPassword(payload.Password)
	if err != nil {
		fmt.Println("Error hashing password in InsertLogin:", err)
		response.Code = http.StatusBadRequest
		response.Message = err.Error()
		return false, response
	}
	model.Password = hashed
	model.UserID = userID

	if result := db.Create(payload); result.Error != nil {
		fmt.Println("Error creating record in InsertLogin:", result.Error)
		response.Code = http.StatusBadRequest
		response.Message = result.Error.Error()
		return false, response
	}
	response.Code = http.StatusCreated
	response.Message = "User Created succesfully"

	return true, response
}

func CheckLogin(c *gin.Context, db *gorm.DB, model Model.Login) (bool, string) {
	payload := &model

	// Retrieve the user from the database
	var user Model.Login
	result := db.Where("username = ?", payload.Username).First(&user)
	if result.Error != nil {
		return false, ""
	}
	// Check the password
	if !Helpers.CheckPassword(payload.Password, user.Password) {
		return false, ""
	}

	if token, err := Helpers.GenerateToken(payload.Username); err != nil {
		return false, ""
	} else {
		return true, token
	}

}
