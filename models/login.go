package model

import (
	"gorm.io/gorm"
)

type Login struct {
	gorm.Model
	UserID   uint   `json:"user_id"`
	Username string `json:"username" gorm:"uniqueIndex"`
	Password string `json:"password"`
}
