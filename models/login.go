package model

import (
	"gorm.io/gorm"
)

type Login struct {
	gorm.Model
	Username string
	Password string
}
