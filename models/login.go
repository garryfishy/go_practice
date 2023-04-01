package model

import (
	"gorm.io/gorm"
)

type Login struct {
	gorm.Model
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
