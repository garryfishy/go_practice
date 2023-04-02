package controllers

import (
	"net/http"

	Services "go_practice/services"
	"go_practice/structs"
	Structs "go_practice/structs"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Hello(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, gin.H{"msg": "hello world"})

}

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

func Login(c *gin.Context, db *gorm.DB) {
	var request structs.Login
	if err := c.BindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	ok, token := Services.CheckLogin(c, db, request.Login)

	if ok {
		c.IndentedJSON(http.StatusAccepted, gin.H{"token": token})
	} else {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
	}
}
