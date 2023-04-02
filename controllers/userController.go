package controllers

import (
	"net/http"

	Services "go_practice/services"
	Structs "go_practice/structs"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(c *gin.Context, db *gorm.DB) {
	var request Structs.Register
	if err := c.BindJSON(&request); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "failed to parse request body"})
		return
	}
	user := request.User
	login := request.Login

	// Insert the User record
	userID, err := Services.InsertUser(c, db, user)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"msg": "Failed to create user"})
		return
	}

	// Insert the Login record with the retrieved UserID
	login.UserID = userID
	if ok := Services.InsertLogin(c, db, login, userID); ok {
		c.IndentedJSON(http.StatusCreated, gin.H{"msg": "User successfully created"})
	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"msg": "Failed to create user"})
	}
}

func Hello(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, gin.H{"msg": "hello world"})

}
