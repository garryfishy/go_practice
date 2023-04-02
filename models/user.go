package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name   string `json:"name"`
	Email  string `json:"email" gorm:"uniqueIndex"`
	Age    int    `json:"age"`
	Logins []Login
}
