package controllers

import (
	"fmt"
	"net/http"

	Services "go_practice/services"
	Structs "go_practice/structs"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Register(c *gin.Context, db *gorm.DB) {
	newUser := Structs.Register{}
	user := Services.InsertUser(c, db, newUser)
	login := Services.InsertLogin(c, db, newUser)
	fmt.Print(user, login, "<<<<< ini")

	if user && login {
		c.IndentedJSON(http.StatusCreated, gin.H{"msg": "User succesfully created"})
	} else {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"msg": "Blah blah bad request"})
	}

}

func Hello(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, gin.H{"msg": "hello world"})

}
