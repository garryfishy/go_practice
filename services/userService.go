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

func InsertUser(c *gin.Context, db *gorm.DB, model Structs.Register) bool {

	payload := &model.User

	fmt.Println(payload, "<<<<payload di user")

	if err := c.BindJSON(payload); err != nil {

		fmt.Println("masuk service insert bindf")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"err": "error di bind"})
		return false
	}

	if result := db.Create(payload); result.Error != nil {
		fmt.Println("masuk service insert creat")
		c.IndentedJSON(http.StatusBadRequest, gin.H{"err": "errir di create"})
		return false
	}
	return true
}

func InsertLogin(c *gin.Context, db *gorm.DB, model Structs.Register) bool {

	var payload = &Model.Login{}
	fmt.Println(payload, "<<<<payload di login")
	if err := c.BindJSON(payload); err != nil {
		// SELALU KENA EOF DI SINI
		c.IndentedJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return false
	}

	hashed, err := Helpers.HashPassword(payload.Password)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return false
	}
	payload.Password = hashed

	if result := db.Create(payload); result.Error != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"err": result.Error.Error()})
		return false
	}
	return true

}
