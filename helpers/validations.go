package helpers

import (
	"errors"
	"strings"

	models "go_practice/models"
)

func ValidateUserModel(user *models.User) error {
	if user.Name == "" {
		return errors.New("name is required")
	}

	if user.Email == "" {
		return errors.New("email is required")
	}

	if len(user.Name) < 2 || len(user.Name) > 50 {
		return errors.New("name length should be between 2 and 50")
	}

	if len(user.Email) < 5 || len(user.Email) > 255 {
		return errors.New("email length should be between 5 and 255")
	}

	if !strings.Contains(user.Email, "@") {
		return errors.New("email should contain @ symbol")
	}

	// add more validations as needed

	return nil
}

func ValidateLogin(login *models.Login) error {
	if len(strings.TrimSpace(login.Username)) == 0 {
		return errors.New("username is required")
	}

	if len(login.Password) < 8 {
		return errors.New("password should be at least 8 characters long")
	}

	hasUppercase := false
	hasLowercase := false
	hasDigit := false

	for _, char := range login.Password {
		if char >= 'A' && char <= 'Z' {
			hasUppercase = true
		} else if char >= 'a' && char <= 'z' {
			hasLowercase = true
		} else if char >= '0' && char <= '9' {
			hasDigit = true
		}
	}

	if !hasUppercase || !hasLowercase || !hasDigit {
		return errors.New("password should contain at least one uppercase letter, one lowercase letter, and one digit")
	}

	return nil
}
