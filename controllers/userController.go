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

	c.JSON(http.StatusOK, gin.H{"msg": "hello world"})

}

func Register(c *gin.Context, db *gorm.DB) {
	var request Structs.Register
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to parse request body"})
		return
	}
	user := request.User
	login := request.Login

	// Insert the User record
	// Start a database transaction
	tx := db.Begin()

	// Insert the User record
	userID, err, Response := Services.InsertUser(c, tx, user)
	if err != nil {
		tx.Rollback() // Roll back the transaction if there's an error
		c.JSON(Response.Code, gin.H{"msg": Response.Message})
		return
	}

	// Update the user object with the newly created UserID
	user.ID = userID

	// Insert the Login record with the retrieved UserID
	if ok, response := Services.InsertLogin(c, tx, login, userID); ok {
		// Commit the transaction if both operations are successful
		tx.Commit()
		c.JSON(response.Code, gin.H{"msg": response.Message})
	} else {
		// Roll back the transaction if there's an error
		tx.Rollback()
		c.JSON(http.StatusBadRequest, gin.H{"msg": "Failed to create user!"})
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
		c.JSON(http.StatusAccepted, gin.H{"token": token})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
	}
}
