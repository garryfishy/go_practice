package structs

import (
	Models "go_practice/models"
)

type Register struct {
	Models.User  `json:"user"`
	Models.Login `json:"login"`
}
