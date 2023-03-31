package structs

import (
	Models "go_practice/models"

	"gorm.io/gorm"
)

type Register struct {
	gorm.Model
	Models.User
	Models.Login
}
